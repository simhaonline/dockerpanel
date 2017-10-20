// This file is automatically generated by qtc from "cuenta.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/cuenta.qtpl:1
package templates

//line templates/cuenta.qtpl:1
import "github.com/mkreder/dockerpanel/model"

//line templates/cuenta.qtpl:2
import "strconv"

//line templates/cuenta.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/cuenta.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/cuenta.qtpl:3
func StreamCuentaTemplate(qw422016 *qt422016.Writer, cuentas []model.Cuenta, dominio model.Dominio, error string) {
	//line templates/cuenta.qtpl:3
	qw422016.N().S(`

<!DOCTYPE html>
<html lang="en">
<script type="text/javascript">
    function validateForm() {
        var y = document.getElementById("nombre").value;
        if (y == "" ) {
            alert("Se debe completar el nombre de la cuenta");
            return false;
        }

        var y = document.getElementById("nombreReal").value;
        if (y == "" ) {
            alert("Se debe completar el nombre real de la cuenta");
            return false;
        }

        if ( document.getElementById('aractivado').checked ) {
            var a = document.getElementById("fechaInicio").value;
            if (a == "" ) {
                alert("Se debe completar la fecha de inicio");
                return false;
            }

            var b = document.getElementById("fechaFin").value;
            if (b == "" ) {
                alert("Se debe completar la fecha de fin");
                return false;
            }

            var d = document.getElementById("asunto").value;
            if (c == "" ) {
                alert("Se debe completar el asunto");
                return false;
            }

            var d = document.getElementById("mensaje").value;
            if (d == "" ) {
                alert("Se debe completar el mensaje");
                return false;
            }
        }

        if ( document.getElementById('renvioactivo').checked ) {
            var e = document.getElementById("direccionRenvio").value;
            if (e == "" ) {
                alert("Se debe completar la dirección de renvío");
                return false;
            }

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


    function showDiv() {
        var x = document.getElementById('form');
        x.style.display = 'block';
    }


    function hideDiv() {
        var x = document.getElementById('form');
        x.style.display = 'none';
        document.getElementById("nombre").value = "";
        document.getElementById("nombreReal").value = "";
        document.getElementById("password").value = "";
        document.getElementById("cuota").value = "";
        document.getElementById("cuentadefecto").checked = false;
        document.getElementById("aractivado").checked = false;
        document.getElementById("fechaInicio").value = "";
        document.getElementById("fechaFin").value = "";
        document.getElementById("mensaje").value = "";
        document.getElementById("asunto").value = "";
        document.getElementById("renvioactivo").checked = false;
        document.getElementById("direccionRenvio").value = "";
        document.getElementById("fechaInicio").setAttribute("disabled","disabled")
        document.getElementById("fechaFin").setAttribute("disabled","disabled")
        document.getElementById("asunto").setAttribute("disabled","disabled")
        document.getElementById("mensaje").setAttribute("disabled","disabled")
        document.getElementById("direccionRenvio").setAttribute("disabled","disabled")

        if ( document.getElementById("password").getAttribute("disabled") == "disabled" ) {
            document.getElementById("btngenerar").innerHTML="Generar";
            document.getElementById("password").removeAttribute("disabled");
            document.getElementById("checkmostrar").removeAttribute("disabled");
        }
    }

    function modifyCuenta(id,nombre,nombreReal,cuota,cuentadefecto,aractivado,fechaInicio,fechaFIn,mensaje,asunto,renvio ){
        var x = document.getElementById('form');
        document.getElementById("id").value = id;
        document.getElementById("nombre").value = nombre;
        document.getElementById("nombreReal").value = nombreReal;
        document.getElementById("cuota").value = cuota;

        if ( cuentadefecto == "si" ) {
            document.getElementById("cuentadefecto").checked = true;
            document.getElementById("direccionRenvio").removeAttribute("disabled")
        } else {
            document.getElementById("cuentadefecto").checked = false;
            document.getElementById("direccionRenvio").setAttribute("disabled","disabled")
        }

        if ( aractivado == "true" ) {
            document.getElementById("aractivado").checked = true;
            document.getElementById("fechaInicio").removeAttribute("disabled")
            document.getElementById("fechaFin").removeAttribute("disabled")
            document.getElementById("asunto").removeAttribute("disabled")
            document.getElementById("mensaje").removeAttribute("disabled")
        } else {
            document.getElementById("aractivado").checked = false;
            document.getElementById("fechaInicio").setAttribute("disabled","disabled")
            document.getElementById("fechaFin").setAttribute("disabled","disabled")
            document.getElementById("asunto").setAttribute("disabled","disabled")
            document.getElementById("mensaje").setAttribute("disabled","disabled")

        }

        document.getElementById("fechaInicio").value = fechaInicio;
        document.getElementById("fechaFin").value = fechaFIn;
        document.getElementById("mensaje").value = mensaje.replace(/(?:;)/g, '\n');
        document.getElementById("asunto").value = asunto;

        if ( renvio == "" ){
            document.getElementById("renvioactivo").checked = false;
        }else {
            document.getElementById("renvioactivo").checked = true;
            document.getElementById("direccionRenvio").value = renvio;
        }


        document.getElementById("password").setAttribute("disabled","disabled")
        document.getElementById("btngenerar").innerHTML="Cambiar";
        document.getElementById("checkmostrar").setAttribute("disabled","disabled")
        x.style.display = 'block';
    }

    function activarRenvio() {
        if (document.getElementById('renvioactivo').checked) {
            document.getElementById("direccionRenvio").removeAttribute("disabled");
        } else {
            document.getElementById("direccionRenvio").setAttribute("disabled","disabled")
            document.getElementById("direccionRenvio").value = "";
        }
    }

    function activarAutoresponder() {
        if (document.getElementById('aractivado').checked) {
            document.getElementById("fechaInicio").removeAttribute("disabled");
            document.getElementById("fechaFin").removeAttribute("disabled");
            document.getElementById("mensaje").removeAttribute("disabled");
            document.getElementById("asunto").removeAttribute("disabled");
        } else {
            document.getElementById("fechaInicio").setAttribute("disabled","disabled")
            document.getElementById("fechaFin").setAttribute("disabled","disabled")
            document.getElementById("mensaje").setAttribute("disabled","disabled")
            document.getElementById("asunto").setAttribute("disabled","disabled")
            document.getElementById("fechaInicio").value = "";
            document.getElementById("fechaFin").value = "";
            document.getElementById("mensaje").value = "";
            document.getElementById("asunto").value = "Fuera de la oficina";
        }
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

<body>

`)
	//line templates/cuenta.qtpl:234
	if len(error) > 0 {
		//line templates/cuenta.qtpl:234
		qw422016.N().S(`
<script type="text/javascript">
    window.alert("`)
		//line templates/cuenta.qtpl:236
		qw422016.E().S(error)
		//line templates/cuenta.qtpl:236
		qw422016.N().S(`")
</script>
`)
		//line templates/cuenta.qtpl:238
	}
	//line templates/cuenta.qtpl:238
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
            <a class="navbar-brand" href="index.html">Docker Panel</a>
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
                        <a href="/"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
                    </li>
                    <li>
                        <a href="/web"><i class="fa fa-server fa-fw"></i>Sitios Web</a>
                    </li>
                    <li>
                        <a href="/dns"><i class="fa fa-cloud fa-fw"></i>DNS</a>
                    </li>
                    <li>
                        <a href="/bd"><i class="fa fa-database fa-fw"></i>Base de Datos</a>
                    </li>
                    <li>
                        <a href="/mail"><i class="fa fa-at fa-fw"></i>E-Mail</a>
                    </li>
                    <li>
                        <a href="/ftp"><i class="fa fa-file-archive-o fa-fw"></i>FTP</a>
                    </li>
                </ul>
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
                    <div class="panel-heading clearfix">
                        <h4 class="panel-title pull-left" style="padding-top: 7.5px;">Cuentas de correo</h4>
                        <button type="button" class="btn pull-right btn-primary btn-sm" onclick="showDiv()">Agregar</button>
                    </div>
                    <!-- /.panel-heading -->
                    <div class="panel-body">
                        <table width="100%" class="table table-striped table-bordered table-hover" id="dataTables-example">
                            <thead>
                            <tr>
                                <th>Cuenta</th>
                                <th>Nombre Real</th>
                                <th>Estado</th>
                                <th>Cuota</th>
                                <th>Auto-responder</th>
                                <th>Cuenta por defecto</th>
                                <th>Renvio de Correo</th>
                                <th>Acciones</th>
                            </tr>
                            </thead>
                            <tbody>
                            `)
	//line templates/cuenta.qtpl:327
	for _, cuenta := range cuentas {
		//line templates/cuenta.qtpl:327
		qw422016.N().S(`
                            <tr class="odd gradeX">
                                <td> `)
		//line templates/cuenta.qtpl:329
		qw422016.E().S(cuenta.Nombre)
		//line templates/cuenta.qtpl:329
		qw422016.N().S(`@`)
		//line templates/cuenta.qtpl:329
		qw422016.E().S(dominio.Nombre)
		//line templates/cuenta.qtpl:329
		qw422016.N().S(` </td>
                                <td> `)
		//line templates/cuenta.qtpl:330
		qw422016.E().S(cuenta.NombreReal)
		//line templates/cuenta.qtpl:330
		qw422016.N().S(` </td>

                                `)
		//line templates/cuenta.qtpl:332
		switch cuenta.Estado {
		//line templates/cuenta.qtpl:333
		case 1:
			//line templates/cuenta.qtpl:333
			qw422016.N().S(`
                                <td>a configurar</td>
                                `)
		//line templates/cuenta.qtpl:335
		case 2:
			//line templates/cuenta.qtpl:335
			qw422016.N().S(`
                                <td>configurado</td>
                                `)
			//line templates/cuenta.qtpl:337
		}
		//line templates/cuenta.qtpl:337
		qw422016.N().S(`

                                `)
		//line templates/cuenta.qtpl:339
		if cuenta.Cuota == 0 {
			//line templates/cuenta.qtpl:339
			qw422016.N().S(`
                                <td> ilimitada </td>
                                `)
			//line templates/cuenta.qtpl:341
		} else {
			//line templates/cuenta.qtpl:341
			qw422016.N().S(`
                                <td> `)
			//line templates/cuenta.qtpl:342
			qw422016.N().D(cuenta.Cuota)
			//line templates/cuenta.qtpl:342
			qw422016.N().S(` MB </td>
                                `)
			//line templates/cuenta.qtpl:343
		}
		//line templates/cuenta.qtpl:343
		qw422016.N().S(`

                                <td>
                                    `)
		//line templates/cuenta.qtpl:346
		if cuenta.Autoresponder.Activado == true {
			//line templates/cuenta.qtpl:346
			qw422016.N().S(`
                                        si
                                    `)
			//line templates/cuenta.qtpl:348
		} else {
			//line templates/cuenta.qtpl:348
			qw422016.N().S(`
                                        no
                                    `)
			//line templates/cuenta.qtpl:350
		}
		//line templates/cuenta.qtpl:350
		qw422016.N().S(`
                                </td>
                                <td>
                                    `)
		//line templates/cuenta.qtpl:354
		cuentadefecto := "no"

		//line templates/cuenta.qtpl:355
		qw422016.N().S(`
                                    `)
		//line templates/cuenta.qtpl:356
		if cuenta.Nombre == dominio.CuentaDefecto {
			//line templates/cuenta.qtpl:356
			qw422016.N().S(`
                                        si
                                    `)
			//line templates/cuenta.qtpl:359
			cuentadefecto = "si"

			//line templates/cuenta.qtpl:360
			qw422016.N().S(`
                                    `)
			//line templates/cuenta.qtpl:361
		} else {
			//line templates/cuenta.qtpl:361
			qw422016.N().S(`
                                        no
                                    `)
			//line templates/cuenta.qtpl:363
		}
		//line templates/cuenta.qtpl:363
		qw422016.N().S(`
                                </td>
                                <td>
                                    `)
		//line templates/cuenta.qtpl:366
		if len(cuenta.Renvio) > 0 {
			//line templates/cuenta.qtpl:366
			qw422016.N().S(`
                                        si
                                    `)
			//line templates/cuenta.qtpl:368
		} else {
			//line templates/cuenta.qtpl:368
			qw422016.N().S(`
                                        no
                                    `)
			//line templates/cuenta.qtpl:370
		}
		//line templates/cuenta.qtpl:370
		qw422016.N().S(`
                                </td>
                                <td class="center">
                                    <button type="button" class="btn btn-xs btn-primary" data-toggle="tooltip" data-placement="top" title="Modificar cuenta de correo" onclick='modifyCuenta(`)
		//line templates/cuenta.qtpl:373
		qw422016.N().D(int(cuenta.ID))
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`, "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(cuenta.Nombre)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`", "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(cuenta.NombreReal)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`","`)
		//line templates/cuenta.qtpl:373
		qw422016.N().D(cuenta.Cuota)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`", "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(cuentadefecto)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`", "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(strconv.FormatBool(cuenta.Autoresponder.Activado))
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`", "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(cuenta.Autoresponder.FechaIncio)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`", "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(cuenta.Autoresponder.FechaFin)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`", "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(cuenta.Autoresponder.Mensaje)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`", "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(cuenta.Autoresponder.Asunto)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`", "`)
		//line templates/cuenta.qtpl:373
		qw422016.E().S(cuenta.Renvio)
		//line templates/cuenta.qtpl:373
		qw422016.N().S(`" )' ><i class="fa fa-list"></i></button>
                                    <button class="btn btn-xs btn-danger" data-toggle="tooltip" data-placement="top" title="Eliminar cuenta de correo" onclick="location.href='removeCuenta?id=`)
		//line templates/cuenta.qtpl:374
		qw422016.N().D(int(cuenta.ID))
		//line templates/cuenta.qtpl:374
		qw422016.N().S(`&dominioid=`)
		//line templates/cuenta.qtpl:374
		qw422016.N().D(int(dominio.ID))
		//line templates/cuenta.qtpl:374
		qw422016.N().S(`';"><i class="fa fa-trash-o"></i></button>
                                </td>
                            </tr>
                            `)
		//line templates/cuenta.qtpl:377
	}
	//line templates/cuenta.qtpl:377
	qw422016.N().S(`
                            </tbody>
                        </table>
                    </div>
                    <!-- /.panel-body -->
                </div>
                <!-- /.panel -->
            </div>
            <!-- /.col-lg-12 -->
        </div>


        <div id="form" class="col-lg-6" hidden="true" >
            <div class="panel panel-default">
                <div class="panel-heading">
                    Configuración de la cuenta de Correo
                </div>
                <div class="panel-body">
                    <form id="addftp" action="/addCuenta" onsubmit="return validateForm()" role=form method="post">

                        <input id="id" name="id" hidden="true" >
                        <input id="dominioid" name="dominioid" hidden="true" value="`)
	//line templates/cuenta.qtpl:398
	qw422016.N().D(int(dominio.ID))
	//line templates/cuenta.qtpl:398
	qw422016.N().S(`">

                        <ul class="nav nav-tabs">
                            <li class="active"><a href="#basica" data-toggle="tab">Configuración Básica</a>
                            </li>
                            <li><a href="#autoresponder" data-toggle="tab">Auto-Responder</a>
                            </li>
                            <li><a href="#renvio" data-toggle="tab">Renvio de Correo</a>
                            </li>
                        </ul>

                        <!-- Tab panes -->
                        <div class="tab-content">
                            <div class="tab-pane fade in active" id="basica">
                                Nombre de la cuenta
                                <input class="form-control" name="nombre" id="nombre">
                                Nombre Real
                                <input class="form-control" name="nombreReal" id="nombreReal">
                                <br>

                                Contraseña
                                <input class="form-control" name="password" id="password" type="password">
                                <button id="btngenerar" type="button" class="btn btn-default btn-sm" onclick="generarPassword()">Generar</button>
                                <input id="checkmostrar" name="checkmostrar" type="checkbox" value="true" onclick="mostrarPassword()"> Mostrar Contraseña
                                <br>

                                <br>
                                Cuota
                                <input id="cuota" type="number" name="cuota" min="0" placeholder="0" value="0" >MB  (0 = ilimitado)
                                <br>

                                <div class="checkbox">
                                    <label>
                                        <input id="cuentadefecto" name="cuentadefecto" type="checkbox" value="true"> Recibir correo de casillas inexistentes
                                    </label>
                                </div>

                            </div>

                            <div class="tab-pane fade" id="autoresponder">
                                <div class="checkbox">
                                    <label>
                                        <input id="aractivado" name="aractivado" type="checkbox" value="true" onclick="activarAutoresponder()"> Activado
                                    </label>
                                </div>
                                Fecha Inicio
                                <input class="form-control" id="fechaInicio" name="fechaInicio" type="date" disabled="disabled">
                                Fecha Fin
                                <input class="form-control" id="fechaFin" name="fechaFin" type="date" disabled="disabled">
                                <br>
                                Asunto
                                <input class="form-control" name="asunto" id="asunto" value="Fuera de la oficina" placeholder="Fuera de la oficina" disabled="disabled">
                                <br>
                                Mensaje
                                <br>
                                <textarea id="mensaje" name="mensaje" cols="50" rows="10" disabled="disabled"></textarea>
                            </div>


                            <div class="tab-pane fade" id="renvio">
                                <div class="checkbox">
                                    <label>
                                        <input id="renvioactivo" name="renvioactivo" type="checkbox" value="true" onclick="activarRenvio()"> Activado
                                    </label>
                                </div>
                                Dirección de renvío
                                <input class="form-control" name="direccionRenvio" id="direccionRenvio" disabled="disabled">

                            </div>



                        </div>
                        <br>
                        <button type="submit" class="btn btn-default">Guardar</button>
                        <button type="button" class="btn btn-default" onclick="hideDiv()">Cancelar</button>
                    </form>
                </div>
            </div>

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
//line templates/cuenta.qtpl:505
}

//line templates/cuenta.qtpl:505
func WriteCuentaTemplate(qq422016 qtio422016.Writer, cuentas []model.Cuenta, dominio model.Dominio, error string) {
	//line templates/cuenta.qtpl:505
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/cuenta.qtpl:505
	StreamCuentaTemplate(qw422016, cuentas, dominio, error)
	//line templates/cuenta.qtpl:505
	qt422016.ReleaseWriter(qw422016)
//line templates/cuenta.qtpl:505
}

//line templates/cuenta.qtpl:505
func CuentaTemplate(cuentas []model.Cuenta, dominio model.Dominio, error string) string {
	//line templates/cuenta.qtpl:505
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/cuenta.qtpl:505
	WriteCuentaTemplate(qb422016, cuentas, dominio, error)
	//line templates/cuenta.qtpl:505
	qs422016 := string(qb422016.B)
	//line templates/cuenta.qtpl:505
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/cuenta.qtpl:505
	return qs422016
//line templates/cuenta.qtpl:505
}