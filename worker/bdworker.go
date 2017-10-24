
package worker

import "log"
import (
	"github.com/mkreder/dockerpanel/model"
	"time"
	"os"
	"os/exec"
	"strconv"
	"io/ioutil"
	"bytes"
)

func calcularPuerto(id uint) string{
	if ( id < 10 ){
		return "300" + strconv.Itoa(int(id))
	} else if (id < 100){
		return "30" + strconv.Itoa(int(id))
	} else if (id < 1000 ){
		return "3" + strconv.Itoa(int(id))
	} else {
		log.Panic("No soportamos mas de 1000 BDs")
	}
	return ""
}

func crearDirectorioConfigBD(nombre string){
	if _ , err := os.Stat("configs/bd/Dockerfile"); os.IsNotExist(err){
		log.Printf("Creando archivos de configuración")
		srcFolder := "defaults/bd"
		os.MkdirAll("configs/bd",0755)
		destFolder := "configs/"
		cpCmd := exec.Command("cp", "-rf", srcFolder, destFolder)
		err := cpCmd.Run()
		check(err)
		os.MkdirAll("data/bd/" + nombre,0755)
	}
}

func crearSQL(bd model.BD){
	sql := ""
	for _ , usuario := range bd.Usuarios {
		for _, ip := range bd.IPs {
			sql = sql + "GRANT ALL ON *.* TO '" + usuario.Nombre + "'@'" + ip.Valor + "' IDENTIFIED BY '" + usuario.Password + "' ;\n"
		}
		// Red de contenedores
		ip := "172.17.0.0/255.255.255.0"
		sql = sql + "GRANT ALL ON *.* TO '" + usuario.Nombre + "'@'" + ip + "' IDENTIFIED BY '" + usuario.Password + "' ;\n"
	}
	sql = sql + "FLUSH PRIVILEGES; \n"
	os.MkdirAll("configs/bd/" + bd.Nombre + "/conf",0755)
	err := ioutil.WriteFile("configs/bd/" + bd.Nombre + "/conf/userdata.sql", []byte(sql), 0644)
	check(err)
}

func construirImagenBD(){
	cmdString := "docker images -q dp-img-bd"
	imgCmd := exec.Command("/bin/sh" , "-c", cmdString)
	var stderr bytes.Buffer
	var out bytes.Buffer
	imgCmd.Stderr = &stderr
	imgCmd.Stdout = &out
	err := imgCmd.Run()
	checkCmd(err,stderr)
	if len(out.String()) == 0 {
		log.Printf("Construyendo la imagen")
		cmdString := "cd configs/bd/" + "; docker build -t dp-img-bd" + " ."
		buildCmd := exec.Command("/bin/sh" , "-c", cmdString)
		var out2 bytes.Buffer
		var stderr2 bytes.Buffer
		buildCmd.Stdout = &out2
		buildCmd.Stderr = &stderr2
		err := buildCmd.Run()
		log.Printf(out2.String())
		checkCmd(err,stderr2)
		log.Printf("Imagen creada")
	} else {
		log.Printf("La imagen ya existe, salteando paso.")
	}
}

func correrContenedorBD(bd model.BD){
	wd, _ := os.Getwd()
	cmdString := "cd configs/bd/" + bd.Nombre + "; docker stop dp-bd-" + bd.Nombre + "; docker rm dp-bd-"+ bd.Nombre +"; docker run -d -p " + calcularPuerto(bd.ID) + ":3306"   +   " -v " + wd + "/configs/bd/" + bd.Nombre + "/conf:/conf:ro"  + " -v " + wd + "/data/bd/" + bd.Nombre + ":/var/lib/mysql" + " --name dp-bd-" + bd.Nombre + " dp-img-bd"
	runCmd := exec.Command("/bin/sh" , "-c", cmdString)
	var stderr bytes.Buffer
	var out bytes.Buffer
	runCmd.Stdout = &out
	runCmd.Stderr = &stderr
	err := runCmd.Run()
	checkCmd(err,stderr)
	log.Printf("Contenedor iniciado dp-bd-" + bd.Nombre + " " + out.String() )
}

func ejecutarSQL(bd model.BD){
	cmdString := "docker exec dp-bd-" + bd.Nombre + " /scripts/loadconf.sh"
	execCmd := exec.Command("/bin/sh" , "-c", cmdString)
	var stderr bytes.Buffer
	var out bytes.Buffer
	execCmd.Stdout = &out
	execCmd.Stderr = &stderr
	err := execCmd.Run()
	checkCmd(err,stderr)
	log.Printf("SQL ejecutado dp-bd-" + bd.Nombre + " " + out.String() )
}
func RunBDWorker() {
	log.Printf("Iniciando BD worker")
	// Loop para siempre
	for {
		for _ , bd := range model.Mgr.GetAllBDs() {
			if bd.Estado == 1 {
				log.Printf("Trabajando en la BD " + bd.Nombre )

				if ! isRunning("dp-bd-" + bd.Nombre){
					crearDirectorioConfigBD(bd.Nombre)
				}
				crearSQL(bd)
				if ! isRunning("dp-bd-" + bd.Nombre) {
					construirImagenBD()
				} else {
					ejecutarSQL(bd)
				}
				bd.Estado = 2
				model.Mgr.UpdateBD(&bd)
			} else if bd.Estado == 2 || bd.Estado == 4 {
				if ! isRunning("dp-bd-" + bd.Nombre){
					correrContenedorBD(bd)
				} else {
					reiniciarContenedor("dp-bd-" + bd.Nombre)
				}
				bd.Estado = 3
				model.Mgr.UpdateBD(&bd)
			} else if bd.Estado == 3 {
				if ! isRunning("dp-bd-" + bd.Nombre){
					bd.Estado = 4
					model.Mgr.UpdateBD(&bd)
				}
			} //else if bd.Estado == 5 {
			//	removeBD(bd)
			//	reiniciarContenedor("dp-bd-" + bd.Nombre)
			//}

		}
		time.Sleep(2 * time.Second)
	}
}