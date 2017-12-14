// This file is automatically generated by qtc from "bd.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/bd.qtpl:1
package templates

//line templates/bd.qtpl:1
import "github.com/mkreder/dockerpanel/model"

//line templates/bd.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/bd.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/bd.qtpl:2
func StreamBDTemplate(qw422016 *qt422016.Writer, bds []model.BD, ubds []model.UsuarioBD, abds []model.AsociacionBD, error string) {
	//line templates/bd.qtpl:2
	qw422016.N().S(`

<!DOCTYPE html>
<html lang="en">
<script src="https://code.jquery.com/jquery-1.11.2.min.js"></script>
<script type="text/javascript">

    $(document).ready(function(){
        $('a[data-toggle="tab"]').on('show.bs.tab', function(e) {
            localStorage.setItem('activeTab', $(e.target).attr('href'));
        });
        var activeTab = localStorage.getItem('activeTab');
        if(activeTab){
            $('#myTab a[href="' + activeTab + '"]').tab('show');
        }
    });

    var timeoutVar;

    function activarTimeout() {
        timeoutVar = setTimeout(function() {
            window.location.href = window.location.href;
        }, 5000);
    }
    activarTimeout()

    function desactivarTimeout() {
        clearTimeout(timeoutVar);
    }

    function validateFormUsuario() {
        var y = document.getElementById("nombre").value;
        if (y == "" ) {
            alert("Se debe completar el nombre de usuario");
            return false;
        }

        var z = document.getElementById("password").value;
        if ( document.getElementById("password").getAttribute("disabled") == "disabled" ){
        } else {
            if (z == "" ) {
                alert("Se debe completar la contraseña");
                return false;
            }
        }

    }

    function generarPassword() {
        if ( document.getElementById("password").getAttribute("disabled") == "disabled" ) {
            document.getElementById("btngenerar").innerHTML="Generar";
            document.getElementById("password").removeAttribute("disabled")
            document.getElementById("checkmostrar").removeAttribute("disabled");
        } else {
            document.getElementById("password").value = Math.random().toString(36).slice(-8);
        }
    }

    function mostrarPassword() {
        if ( document.getElementById('checkmostrar').checked ) {
            document.getElementById('password').removeAttribute("type")
        } else {
            document.getElementById('password').setAttribute("type","password")
        }

    }


    function showDivUsuario() {
        var x = document.getElementById('formUsuario');
        x.style.display = 'block';
        desactivarTimeout()
    }


    function hideDivUsuario() {
        var x = document.getElementById('formUsuario');
        x.style.display = 'none';
        document.getElementById("nombre").value = "";
        document.getElementById("password").value = "";
        if ( document.getElementById("password").getAttribute("disabled") == "disabled" ) {
            document.getElementById("btngenerar").innerHTML="Generar";
            document.getElementById("password").removeAttribute("disabled");
            document.getElementById("checkmostrar").removeAttribute("disabled");
        }
        activarTimeout()
    }

    function modifyUsuario(id,nombre) {
        var x = document.getElementById('formUsuario');
        document.getElementById("id").value = id;
        document.getElementById("nombre").value = nombre;
        document.getElementById("password").setAttribute("disabled","disabled")
        document.getElementById("btngenerar").innerHTML="Cambiar";
        document.getElementById("checkmostrar").setAttribute("disabled","disabled")

        x.style.display = 'block';
        desactivarTimeout()
    }


    function showDivBD() {
        var x = document.getElementById('formBD');
        x.style.display = 'block';
        desactivarTimeout()
    }

    function hideDivBD() {
        var x = document.getElementById('formBD');
        x.style.display = 'none';
        document.getElementById("nombre").value = "";
        activarTimeout()
    }


    function hideDivRemoveBD(id) {
        var x = document.getElementById('formRemoveUBD' + id);
        x.style.display = 'none';
        document.getElementById("bd" + id).options[0].selected = true;
        activarTimeout()
    }

    function hideDivAUBD() {
        var x = document.getElementById('formAddUBD');
        x.style.display = 'none';
        document.getElementById("selabd").options[0].selected = true;
        activarTimeout()
    }

    function validateFormBD() {
        var y = document.getElementById("nombreBD").value;
        if (y == "" ) {
            alert("Se debe completar el nombre de la base de datos");
            return false;
        }
    }

    function validateFormAddUBD() {
        var y = document.getElementById("selabd").value;
        if (y == "" ) {
            alert("Se debe completar el nombre de la base de datos");
            return false;
        }
    }

    function validateFormRemoveUBD(id) {
        var y = document.getElementById("bd" + id).value;
        if (y == "" ) {
            alert("Se debe seleccionar una base de datos");
            return false;
        }
    }

    function validateFormRemoveIP(id) {
        var y = document.getElementById("ip" + id).value;
        if (y == "" ) {
            alert("Se debe seleccionar una IP");
            return false;
        }
    }


    function modifyBD(id,nombre) {
        var x = document.getElementById('formBD');
        document.getElementById("idBD").value = id;
        document.getElementById("nombreBD").value = nombre;
        x.style.display = 'block';
        desactivarTimeout()

    }

    function addUBD(id){
        var x = document.getElementById('formAddUBD');
        document.getElementById("idABD").value = id;
        x.style.display = 'block';
        desactivarTimeout()
    }

    function removeUBD(id){
        var x = document.getElementById('formRemoveUBD' + id);
        x.style.display = 'block';
        desactivarTimeout()
    }

    function addIP(id){
        var x = document.getElementById('formAddIP');
        document.getElementById("idIP").value = id;
        x.style.display = 'block';
        desactivarTimeout()
    }

    function removeIP(id){
        var x = document.getElementById('formRemoveIP' + id);
        x.style.display = 'block';
        desactivarTimeout()
    }

    function validateFormAddIP(){
        var y = document.getElementById("ip").value;
        if (y == "" ) {
            alert("Se debe completar la IP");
            return false;
        }

        if (/^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^%$/.test(y)) {
        }else {
            alert("IP invalida");
            return false;
        }
    }

    function hideDivAIP(){
        var x = document.getElementById('formAddIP');
        x.style.display = 'none';
        document.getElementById("ip").value = "";
        activarTimeout()
    }

    function hideDivRemoveIP(id) {
        var x = document.getElementById('formRemoveIP' + id);
        x.style.display = 'none';
        document.getElementById("ip" + id).options[0].selected = true;
        activarTimeout()
    }

    function setHref() {
        document.getElementById('phpmyadmin').href = window.location.protocol + "//" + window.location.hostname + ":58080/";
    }




</script>
<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Panel Docker</title>

    <!-- Bootstrap Core CSS -->
    <link href="../vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">

    <!-- MetisMenu CSS -->
    <link href="../vendor/metisMenu/metisMenu.min.css" rel="stylesheet">

    <!-- Custom CSS -->
    <link href="../dist/css/sb-admin-2.css" rel="stylesheet">

    <!-- Morris Charts CSS -->
    <link href="../vendor/morrisjs/morris.css" rel="stylesheet">

    <!-- Custom Fonts -->
    <link href="../vendor/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
    <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->



</head>

<body onload="setHref()">>

`)
	//line templates/bd.qtpl:273
	if len(error) > 0 {
		//line templates/bd.qtpl:273
		qw422016.N().S(`
<script type="text/javascript">
    window.alert("`)
		//line templates/bd.qtpl:275
		qw422016.E().S(error)
		//line templates/bd.qtpl:275
		qw422016.N().S(`")


</script>
`)
		//line templates/bd.qtpl:279
	}
	//line templates/bd.qtpl:279
	qw422016.N().S(`

<div id="wrapper">

    <!-- Navigation -->
    <nav class="navbar navbar-default navbar-static-top" role="navigation" style="margin-bottom: 0">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand">Docker Panel</a>
        </div>
        <!-- /.navbar-header -->

        <ul class="nav navbar-top-links navbar-right">
            <!-- /.dropdown -->
            <li class="dropdown">
                <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                    <i class="fa fa-user fa-fw"></i> <i class="fa fa-caret-down"></i>
                </a>
                <ul class="dropdown-menu dropdown-user">
                    <li><a href="/profile"><i class="fa fa-user fa-fw"></i> Configuración</a>
                    </li>
                    <li class="divider"></li>
                    <li><a href="/logout"><i class="fa fa-sign-out fa-fw"></i> Logout</a>
                    </li>
                </ul>
                <!-- /.dropdown-user -->
            </li>
            <!-- /.dropdown -->
        </ul>
        <!-- /.navbar-top-links -->


        <div class="navbar-default sidebar" role="navigation">
            <div class="sidebar-nav navbar-collapse">
                <ul class="nav" id="side-menu">
                    <li>
                        <a href="/"><i class="fa fa-dashboard fa-fw"></i>Principal</a>
                    </li>
                    <li>
                        <a href="/web"><i class="fa fa-server fa-fw"></i>Sitios web</a>
                    </li>
                    <li>
                        <a href="/dns"><i class="fa fa-cloud fa-fw"></i>DNS</a>
                    </li>
                    <li>
                        <a href="/bd"><i class="fa fa-database fa-fw"></i>Bases de datos</a>
                    </li>
                    <li>
                        <a href="/mail"><i class="fa fa-at fa-fw"></i>E-Mail</a>
                    </li>
                    <li>
                        <a href="/ftp"><i class="fa fa-file-archive-o fa-fw"></i>Usuarios FTP</a>
                    </li>
                </ul>
            </div>
            <!-- /.sidebar-collapse -->
        </div>
            <!-- /.sidebar-collapse -->
        </div>
        <!-- /.navbar-static-side -->
    </nav>

    <div id="page-wrapper">
        <br>
        <div class="row">
            <div class="col-lg-12">


                <div class="panel panel-default">
                    <div class="panel-heading">
                        Bases de datos
                    </div>
                    <!-- /.panel-heading -->
                    <div class="panel-body">
                        <!-- Nav tabs -->
                        <ul class="nav nav-tabs" id="myTab">
                            <li class="active"><a href="#bds" data-toggle="tab">Bases de datos</a>
                            </li>
                            <li><a href="#usuarios" data-toggle="tab">Usuarios</a>
                            </li>
                        </ul>

                        <!-- Tab panes -->
                        <div class="tab-content">
                            <div class="tab-pane fade" id="usuarios">
                                <br>
                                <table width="100%" class="table table-striped table-bordered table-hover" id="dataTables-example">
                                    <thead>
                                    <tr>
                                        <th>Nombre</th>
                                        <th>Bases</th>
                                        <th>Acciones</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    `)
	//line templates/bd.qtpl:379
	for _, user := range ubds {
		//line templates/bd.qtpl:379
		qw422016.N().S(`
                                    <tr class="odd gradeX">
                                        <td> `)
		//line templates/bd.qtpl:381
		qw422016.E().S(user.Nombre)
		//line templates/bd.qtpl:381
		qw422016.N().S(` </td>

                                        <td>
                                            `)
		//line templates/bd.qtpl:384
		for _, abd := range abds {
			//line templates/bd.qtpl:384
			qw422016.N().S(`
                                                `)
			//line templates/bd.qtpl:385
			if user.ID == abd.UsuarioBDID {
				//line templates/bd.qtpl:385
				qw422016.N().S(`
                                                    `)
				//line templates/bd.qtpl:386
				for _, bd := range bds {
					//line templates/bd.qtpl:386
					qw422016.N().S(`
                                                        `)
					//line templates/bd.qtpl:387
					if abd.BDID == bd.ID {
						//line templates/bd.qtpl:387
						qw422016.N().S(`
                                                            `)
						//line templates/bd.qtpl:388
						qw422016.E().S(bd.Nombre)
						//line templates/bd.qtpl:388
						qw422016.N().S(`
                                                        `)
						//line templates/bd.qtpl:389
					}
					//line templates/bd.qtpl:389
					qw422016.N().S(`
                                                    `)
					//line templates/bd.qtpl:390
				}
				//line templates/bd.qtpl:390
				qw422016.N().S(`
                                                `)
				//line templates/bd.qtpl:391
			}
			//line templates/bd.qtpl:391
			qw422016.N().S(`
                                            `)
			//line templates/bd.qtpl:392
		}
		//line templates/bd.qtpl:392
		qw422016.N().S(`
                                        </td>


                                        <td class="center">
                                            <button type="button" class="btn btn-xs btn-success" data-toggle="tooltip" data-placement="top" title="Asociar una base de datos" onclick='addUBD(`)
		//line templates/bd.qtpl:397
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:397
		qw422016.N().S(`)' ><i class="fa fa-plus"></i></button>
                                            <button type="button" class="btn btn-xs btn-warning" data-toggle="tooltip" data-placement="top" title="Desacociar una base de datos" onclick='removeUBD(`)
		//line templates/bd.qtpl:398
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:398
		qw422016.N().S(`)' ><i class="fa fa-minus"></i></button>
                                            <button type="button" class="btn btn-xs btn-primary" data-toggle="tooltip" data-placement="top" title="Editar usuario" onclick='modifyUsuario(`)
		//line templates/bd.qtpl:399
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:399
		qw422016.N().S(`, "`)
		//line templates/bd.qtpl:399
		qw422016.E().S(user.Nombre)
		//line templates/bd.qtpl:399
		qw422016.N().S(`"  )' ><i class="fa fa-list"></i></button>
                                            <button class="btn btn-xs btn-danger" data-toggle="tooltip" data-placement="top" title="Borrar usuario" onclick="location.href='removeUsuarioBD?id=`)
		//line templates/bd.qtpl:400
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:400
		qw422016.N().S(`';"><i class="fa fa-trash-o"></i></button>
                                        </td>
                                    </tr>
                                    `)
		//line templates/bd.qtpl:403
	}
	//line templates/bd.qtpl:403
	qw422016.N().S(`
                                    </tbody>
                                </table>
                                <button type="button" class="btn btn-primary btn-sm" onclick="showDivUsuario()">Agregar</button>
                            </div>
                            <div class="tab-pane fade in active" id="bds">
                                <br>
                                <table width="100%" class="table table-striped table-bordered table-hover" id="dataTables-example">
                                    <thead>
                                    <tr>
                                        <th>Nombre</th>
                                        <th>Estado</th>
                                        <th>Puerto</th>
                                        <th>Remota</th>
                                        <th>Acciones</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    `)
	//line templates/bd.qtpl:421
	for _, bd := range bds {
		//line templates/bd.qtpl:421
		qw422016.N().S(`
                                    <tr class="odd gradeX">
                                        <td> `)
		//line templates/bd.qtpl:423
		qw422016.E().S(bd.Nombre)
		//line templates/bd.qtpl:423
		qw422016.N().S(` </td>


                                        `)
		//line templates/bd.qtpl:426
		switch bd.Estado {
		//line templates/bd.qtpl:427
		case 1:
			//line templates/bd.qtpl:427
			qw422016.N().S(`
                                        <td>a configurar</td>
                                        `)
		//line templates/bd.qtpl:429
		case 2:
			//line templates/bd.qtpl:429
			qw422016.N().S(`
                                        <td>configurado</td>
                                        `)
		//line templates/bd.qtpl:431
		case 3:
			//line templates/bd.qtpl:431
			qw422016.N().S(`
                                        <td>activo</td>
                                        `)
		//line templates/bd.qtpl:433
		case 4:
			//line templates/bd.qtpl:433
			qw422016.N().S(`
                                        <td>apagado</td>
                                        `)
		//line templates/bd.qtpl:435
		case 5:
			//line templates/bd.qtpl:435
			qw422016.N().S(`
                                        <td>eliminando</td>
                                        `)
			//line templates/bd.qtpl:437
		}
		//line templates/bd.qtpl:437
		qw422016.N().S(`

                                        <td>
                                            `)
		//line templates/bd.qtpl:440
		if bd.ID < 10 {
			//line templates/bd.qtpl:440
			qw422016.N().S(`
                                                300`)
			//line templates/bd.qtpl:441
			qw422016.N().D(int(bd.ID))
			//line templates/bd.qtpl:441
			qw422016.N().S(`
                                            `)
			//line templates/bd.qtpl:442
		} else if bd.ID < 100 {
			//line templates/bd.qtpl:442
			qw422016.N().S(`
                                                30`)
			//line templates/bd.qtpl:443
			qw422016.N().D(int(bd.ID))
			//line templates/bd.qtpl:443
			qw422016.N().S(`
                                            `)
			//line templates/bd.qtpl:444
		} else {
			//line templates/bd.qtpl:444
			qw422016.N().S(`
                                                3`)
			//line templates/bd.qtpl:445
			qw422016.N().D(int(bd.ID))
			//line templates/bd.qtpl:445
			qw422016.N().S(`
                                            `)
			//line templates/bd.qtpl:446
		}
		//line templates/bd.qtpl:446
		qw422016.N().S(`

                                        </td>
                                        <td>
                                            `)
		//line templates/bd.qtpl:450
		if len(bd.IPs) == 0 {
			//line templates/bd.qtpl:450
			qw422016.N().S(`
                                                no
                                            `)
			//line templates/bd.qtpl:452
		} else {
			//line templates/bd.qtpl:452
			qw422016.N().S(`
                                                `)
			//line templates/bd.qtpl:453
			for _, ip := range bd.IPs {
				//line templates/bd.qtpl:453
				qw422016.N().S(`
                                                    `)
				//line templates/bd.qtpl:454
				qw422016.E().S(ip.Valor)
				//line templates/bd.qtpl:454
				qw422016.N().S(`
                                                `)
				//line templates/bd.qtpl:455
			}
			//line templates/bd.qtpl:455
			qw422016.N().S(`
                                            `)
			//line templates/bd.qtpl:456
		}
		//line templates/bd.qtpl:456
		qw422016.N().S(`
                                        </td>

                                        <td class="center">
                                            <button type="button" class="btn btn-xs btn-success" data-toggle="tooltip" data-placement="top" title="Agregar IP remota" onclick='addIP(`)
		//line templates/bd.qtpl:460
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:460
		qw422016.N().S(`)' ><i class="fa fa-plus"></i></button>
                                            <button type="button" class="btn btn-xs btn-warning" data-toggle="tooltip" data-placement="top" title="Borrar IP remota" onclick='removeIP(`)
		//line templates/bd.qtpl:461
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:461
		qw422016.N().S(`)' ><i class="fa fa-minus"></i></button>
                                            <button type="button" class="btn btn-xs btn-primary" data-toggle="tooltip" data-placement="top" title="Editar base de datos" onclick='modifyBD(`)
		//line templates/bd.qtpl:462
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:462
		qw422016.N().S(`, "`)
		//line templates/bd.qtpl:462
		qw422016.E().S(bd.Nombre)
		//line templates/bd.qtpl:462
		qw422016.N().S(`")' ><i class="fa fa-list"></i></button>
                                           <button class="btn btn-xs btn-danger"data-toggle="tooltip" data-placement="top" title="Borrar base de datos"  onclick="location.href='removeBd?id=`)
		//line templates/bd.qtpl:463
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:463
		qw422016.N().S(`';"><i class="fa fa-trash-o"></i></button>
                                        </td>
                                    </tr>
                                    `)
		//line templates/bd.qtpl:466
	}
	//line templates/bd.qtpl:466
	qw422016.N().S(`
                                    </tbody>
                                </table>
                                <button type="button" class="btn btn-primary btn-sm" onclick="showDivBD()">Agregar</button>
                            </div>
                        </div>
                        <br>
                        <br>
                        <a class="btn btn-info btn-sm" href="/" id="phpmyadmin">PHPMyAdmin</a>
                    </div>
                    <!-- /.panel-body -->
                </div>





                </div>
                <!-- /.panel -->
            </div>
            <!-- /.col-lg-12 -->



        <div id="formUsuario" class="col-lg-6" hidden="true" >
            <div class="panel panel-default">
                <div class="panel-heading">
                    Configuración del usuario
                </div>
                <div class="panel-body">
                    <form id="addubd" action="/ubd" onsubmit="return validateFormUsuario()" role=form method="post">
                        <input id="id" name="id" hidden="true" >
                        Nombre de usuario
                        <input class="form-control" name="nombre" id="nombre">
                        <br>

                        Contraseña
                        <input class="form-control" name="password" id="password" type="password">
                        <button id="btngenerar" type="button" class="btn btn-default btn-sm" onclick="generarPassword()">Generar</button>
                        <input id="checkmostrar" name="checkmostrar" type="checkbox" value="true" onclick="mostrarPassword()"> Mostrar contraseña
                        <br>

                        <br>
                        <button type="submit" class="btn btn-default">Guardar</button>
                        <button type="button" class="btn btn-default" onclick="hideDivUsuario()">Cancelar</button>
                    </form>
                </div>
            </div>

        </div>



        <div id="formBD" class="col-lg-6" hidden="true" >
            <div class="panel panel-default">
                <div class="panel-heading">
                    Configuración de la base de datos
                </div>
                <div class="panel-body">
                    <form id="addbd" action="/bd" onsubmit="return validateFormBD()" role=form method="post">
                        <input id="idBD" name="idBD" hidden="true" >
                        Nombre de la base de datos
                        <input class="form-control" name="nombreBD" id="nombreBD">
                        <br>
                        <button type="submit" class="btn btn-default">Guardar</button>
                        <button type="button" class="btn btn-default" onclick="hideDivBD()">Cancelar</button>
                    </form>
                </div>
            </div>
        </div>

        <div id="formAddUBD" class="col-lg-6" hidden="true" >
            <div class="panel panel-default">
                <div class="panel-heading">
                    Asociar usuario a base de datos
                </div>
                <div class="panel-body">
                    <form id="adduubd" action="/addubd" onsubmit="return validateFormAddUBD()" role=form method="post">
                        <input id="idABD" name="userid" hidden="true" >
                        Base de datos
                        <select  class="form-control"  id="selabd" name="bdid">
                            <option disabled selected value> -- Elegir base de datos -- </option>
                            `)
	//line templates/bd.qtpl:548
	for _, bd := range bds {
		//line templates/bd.qtpl:548
		qw422016.N().S(`
                            <option value="`)
		//line templates/bd.qtpl:549
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:549
		qw422016.N().S(`">`)
		//line templates/bd.qtpl:549
		qw422016.E().S(bd.Nombre)
		//line templates/bd.qtpl:549
		qw422016.N().S(`</option>
                            `)
		//line templates/bd.qtpl:550
	}
	//line templates/bd.qtpl:550
	qw422016.N().S(`
                        </select>
                        <br>
                        <button type="submit" class="btn btn-default">Guardar</button>
                        <button type="button" class="btn btn-default" onclick="hideDivAUBD()">Cancelar</button>
                    </form>
                </div>
            </div>
        </div>

        `)
	//line templates/bd.qtpl:560
	for _, user := range ubds {
		//line templates/bd.qtpl:560
		qw422016.N().S(`
        <div id="formRemoveUBD`)
		//line templates/bd.qtpl:561
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:561
		qw422016.N().S(`" class="col-lg-6" hidden="true" >
            <div class="panel panel-default">
                <div class="panel-heading">
                    Remover permisos de usuario de base de datos
                </div>
                <div class="panel-body">
                    <form id="removeubd" action="/removeubd" onsubmit="return validateFormRemoveUBD(`)
		//line templates/bd.qtpl:567
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:567
		qw422016.N().S(`)" role=form method="post">
                        <input id="userid" name="userid" value="`)
		//line templates/bd.qtpl:568
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:568
		qw422016.N().S(`" hidden="true" >
                        Base de datos
                        <select  class="form-control"  id="bd`)
		//line templates/bd.qtpl:570
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:570
		qw422016.N().S(`" name="bdid">
                            <option disabled selected value> -- Elegir base de datos -- </option>
                            `)
		//line templates/bd.qtpl:572
		for _, abd := range abds {
			//line templates/bd.qtpl:572
			qw422016.N().S(`
                                `)
			//line templates/bd.qtpl:573
			if user.ID == abd.UsuarioBDID {
				//line templates/bd.qtpl:573
				qw422016.N().S(`
                                    `)
				//line templates/bd.qtpl:574
				for _, bd := range bds {
					//line templates/bd.qtpl:574
					qw422016.N().S(`
                                        `)
					//line templates/bd.qtpl:575
					if abd.BDID == bd.ID {
						//line templates/bd.qtpl:575
						qw422016.N().S(`
                                            <option value="`)
						//line templates/bd.qtpl:576
						qw422016.N().D(int(bd.ID))
						//line templates/bd.qtpl:576
						qw422016.N().S(`">`)
						//line templates/bd.qtpl:576
						qw422016.E().S(bd.Nombre)
						//line templates/bd.qtpl:576
						qw422016.N().S(`</option>
                                        `)
						//line templates/bd.qtpl:577
					}
					//line templates/bd.qtpl:577
					qw422016.N().S(`
                                    `)
					//line templates/bd.qtpl:578
				}
				//line templates/bd.qtpl:578
				qw422016.N().S(`
                                `)
				//line templates/bd.qtpl:579
			}
			//line templates/bd.qtpl:579
			qw422016.N().S(`
                            `)
			//line templates/bd.qtpl:580
		}
		//line templates/bd.qtpl:580
		qw422016.N().S(`
                        </select>
                        <br>
                        <button type="submit" class="btn btn-default">Guardar</button>
                        <button type="button" class="btn btn-default" onclick="hideDivRemoveBD(`)
		//line templates/bd.qtpl:584
		qw422016.N().D(int(user.ID))
		//line templates/bd.qtpl:584
		qw422016.N().S(`)">Cancelar</button>
                    </form>
                </div>
            </div>
        </div>
        `)
		//line templates/bd.qtpl:589
	}
	//line templates/bd.qtpl:589
	qw422016.N().S(`

        <div id="formAddIP" class="col-lg-6" hidden="true" >
            <div class="panel panel-default">
                <div class="panel-heading">
                    Agregar IP remota
                </div>
                <div class="panel-body">
                    <form id="addbdip" action="/addbdip" onsubmit="return validateFormAddIP()" role=form method="post">
                        <input id="idIP" name="bdid" hidden="true" >
                        IP
                        <input id="ip" name="ip" class="form-control" >
                        <br>
                        <button type="submit" class="btn btn-default">Guardar</button>
                        <button type="button" class="btn btn-default" onclick="hideDivAIP()">Cancelar</button>
                    </form>
                </div>
            </div>
        </div>


        `)
	//line templates/bd.qtpl:610
	for _, bd := range bds {
		//line templates/bd.qtpl:610
		qw422016.N().S(`
        <div id="formRemoveIP`)
		//line templates/bd.qtpl:611
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:611
		qw422016.N().S(`" class="col-lg-6" hidden="true" >
            <div class="panel panel-default">
                <div class="panel-heading">
                    Desasociar IP remota
                </div>
                <div class="panel-body">
                    <form id="removebdip" action="/removebdip" onsubmit="return validateFormRemoveIP(`)
		//line templates/bd.qtpl:617
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:617
		qw422016.N().S(`)" role=form method="post">
                        <input id="bd`)
		//line templates/bd.qtpl:618
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:618
		qw422016.N().S(`" name="bdid" value="`)
		//line templates/bd.qtpl:618
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:618
		qw422016.N().S(`" hidden="true" >
                        IPs
                        <select  class="form-control"  id="ip`)
		//line templates/bd.qtpl:620
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:620
		qw422016.N().S(`" name="ip">
                            <option disabled selected value> -- Elegir IP -- </option>
                            `)
		//line templates/bd.qtpl:622
		for _, ip := range bd.IPs {
			//line templates/bd.qtpl:622
			qw422016.N().S(`
                            <option value="`)
			//line templates/bd.qtpl:623
			qw422016.E().S(ip.Valor)
			//line templates/bd.qtpl:623
			qw422016.N().S(`">`)
			//line templates/bd.qtpl:623
			qw422016.E().S(ip.Valor)
			//line templates/bd.qtpl:623
			qw422016.N().S(`</option>
                            `)
			//line templates/bd.qtpl:624
		}
		//line templates/bd.qtpl:624
		qw422016.N().S(`
                        </select>
                        <br>
                        <button type="submit" class="btn btn-default">Guardar</button>
                        <button type="button" class="btn btn-default" onclick="hideDivRemoveIP(`)
		//line templates/bd.qtpl:628
		qw422016.N().D(int(bd.ID))
		//line templates/bd.qtpl:628
		qw422016.N().S(`)">Cancelar</button>
                    </form>
                </div>
            </div>
        </div>
        `)
		//line templates/bd.qtpl:633
	}
	//line templates/bd.qtpl:633
	qw422016.N().S(`


    </div>
        <!-- /#page-wrapper -->

    </div>
    <!-- /#wrapper -->

    <!-- jQuery -->
    <script src="../vendor/jquery/jquery.min.js"></script>

    <!-- Bootstrap Core JavaScript -->
    <script src="../vendor/bootstrap/js/bootstrap.min.js"></script>

    <!-- Metis Menu Plugin JavaScript -->
    <script src="../vendor/metisMenu/metisMenu.min.js"></script>

    <!-- Morris Charts JavaScript -->
    <!--<script src="../vendor/raphael/raphael.min.js"></script>-->
    <!--<script src="../vendor/morrisjs/morris.min.js"></script>-->
    <!--<script src="../data/morris-data.js"></script>-->

    <!-- Custom Theme JavaScript -->
    <script src="../dist/js/sb-admin-2.js"></script>

</body>

</html>

`)
//line templates/bd.qtpl:663
}

//line templates/bd.qtpl:663
func WriteBDTemplate(qq422016 qtio422016.Writer, bds []model.BD, ubds []model.UsuarioBD, abds []model.AsociacionBD, error string) {
	//line templates/bd.qtpl:663
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/bd.qtpl:663
	StreamBDTemplate(qw422016, bds, ubds, abds, error)
	//line templates/bd.qtpl:663
	qt422016.ReleaseWriter(qw422016)
//line templates/bd.qtpl:663
}

//line templates/bd.qtpl:663
func BDTemplate(bds []model.BD, ubds []model.UsuarioBD, abds []model.AsociacionBD, error string) string {
	//line templates/bd.qtpl:663
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/bd.qtpl:663
	WriteBDTemplate(qb422016, bds, ubds, abds, error)
	//line templates/bd.qtpl:663
	qs422016 := string(qb422016.B)
	//line templates/bd.qtpl:663
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/bd.qtpl:663
	return qs422016
//line templates/bd.qtpl:663
}
