// This file is automatically generated by qtc from "mail.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/mail.qtpl:1
package templates

//line templates/mail.qtpl:1
import "github.com/mkreder/dockerpanel/model"

//line templates/mail.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/mail.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/mail.qtpl:2
func StreamMailTemplate(qw422016 *qt422016.Writer, dominios []model.Dominio, error string) {
	//line templates/mail.qtpl:2
	qw422016.N().S(`

<!DOCTYPE html>
<html lang="en">
<script type="text/javascript">
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

    function validateForm() {
        var x = document.getElementById("nombre").value;
        if (x == "" ) {
            alert("Se debe completar el nombre de dominio");
            return false;
        }

        if (/^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9](?:\.[a-zA-Z]{2,})+$/.test(x)) {
        } else {
            alert("Nombre de dominio invalido");
            return false;
        }

        var y = document.getElementById("filtro").value;
        if (y == "" ) {
            alert("Se debe elegir un filtro");
            return false;
        }
    }

    function showDiv() {
        var x = document.getElementById('form');
        x.style.display = 'block';
        desactivarTimeout()
    }


    function hideDiv() {
        var x = document.getElementById('form');
        x.style.display = 'none';
        document.getElementById("nombre").value = "";
        document.getElementById("filtro").options[0].selected = true;
        document.getElementById("id").value = "";
        activarTimeout()
    }

    function modifyDominio(id,nombre,filtro) {
        var x = document.getElementById('form');
        document.getElementById("id").value = id;
        document.getElementById("nombre").value = nombre;

        var selFiltro = document.getElementById("filtro");

        for (var i = 0; i < selFiltro.options.length; i++) {
            if (selFiltro.options[i].value == filtro) {
                selFiltro.options[i].selected = true;
            }
        }

        x.style.display = 'block';
        desactivarTimeout()
    }

    function setHref() {
        document.getElementById('webmail').href = window.location.protocol + "//" + window.location.hostname + ":9080/";
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

<body onload="setHref()">

`)
	//line templates/mail.qtpl:114
	if len(error) > 0 {
		//line templates/mail.qtpl:114
		qw422016.N().S(`
<script type="text/javascript">
    window.alert("`)
		//line templates/mail.qtpl:116
		qw422016.E().S(error)
		//line templates/mail.qtpl:116
		qw422016.N().S(`")
</script>
`)
		//line templates/mail.qtpl:118
	}
	//line templates/mail.qtpl:118
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
            <a class="navbar-brand" >Docker Panel</a>
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
                    <div class="panel-heading clearfix">
                        <h4 class="panel-title pull-left" style="padding-top: 7.5px;">Dominios de correo electrónico</h4>
                        <button type="button" class="btn pull-right btn-primary btn-sm btn-space" onclick="showDiv()">Agregar</button>
                    </div>
                    <!-- /.panel-heading -->
                    <div class="panel-body">
                        <table width="100%" class="table table-striped table-bordered table-hover" id="dataTables-example">
                            <thead>
                            <tr>
                                <th>Dominio</th>
                                <th>Estado</th>
                                <th>Filtro correo basura</th>
                                <th>Acciones</th>
                            </tr>
                            </thead>
                            <tbody>
                            `)
	//line templates/mail.qtpl:205
	for _, dominio := range dominios {
		//line templates/mail.qtpl:205
		qw422016.N().S(`
                            <tr class="odd gradeX">
                                <td> `)
		//line templates/mail.qtpl:207
		qw422016.E().S(dominio.Nombre)
		//line templates/mail.qtpl:207
		qw422016.N().S(` </td>

                                `)
		//line templates/mail.qtpl:209
		switch dominio.Estado {
		//line templates/mail.qtpl:210
		case 1:
			//line templates/mail.qtpl:210
			qw422016.N().S(`
                                <td>a configurar</td>
                                `)
		//line templates/mail.qtpl:212
		case 2:
			//line templates/mail.qtpl:212
			qw422016.N().S(`
                                <td>configurado</td>
                                `)
		//line templates/mail.qtpl:214
		case 3:
			//line templates/mail.qtpl:214
			qw422016.N().S(`
                                <td>activo</td>
                                `)
		//line templates/mail.qtpl:216
		case 4:
			//line templates/mail.qtpl:216
			qw422016.N().S(`
                                <td>apagado</td>
                                `)
		//line templates/mail.qtpl:218
		case 5:
			//line templates/mail.qtpl:218
			qw422016.N().S(`
                                <td>eliminando</td>
                                `)
			//line templates/mail.qtpl:220
		}
		//line templates/mail.qtpl:220
		qw422016.N().S(`

                                <td> `)
		//line templates/mail.qtpl:222
		qw422016.E().S(dominio.FiltroSpam)
		//line templates/mail.qtpl:222
		qw422016.N().S(` </td>

                                <td class="center">
                                    <button class="btn btn-xs btn-success"  data-toggle="tooltip" data-placement="top" title="Editar cuentas" onclick="location.href='editCuentas?dominioid=`)
		//line templates/mail.qtpl:225
		qw422016.N().D(int(dominio.ID))
		//line templates/mail.qtpl:225
		qw422016.N().S(`';"><i class="fa fa-user"></i></button>
                                    <button class="btn btn-xs btn-info"  data-toggle="tooltip" data-placement="top" title="Editar listas de correo" onclick="location.href='editListas?dominioid=`)
		//line templates/mail.qtpl:226
		qw422016.N().D(int(dominio.ID))
		//line templates/mail.qtpl:226
		qw422016.N().S(`';"><i class="fa fa-users"></i></button>
                                    <button type="button" class="btn btn-xs btn-primary" data-toggle="tooltip" data-placement="top" title="Editar dominio" onclick='modifyDominio(`)
		//line templates/mail.qtpl:227
		qw422016.N().D(int(dominio.ID))
		//line templates/mail.qtpl:227
		qw422016.N().S(`, "`)
		//line templates/mail.qtpl:227
		qw422016.E().S(dominio.Nombre)
		//line templates/mail.qtpl:227
		qw422016.N().S(`",  "`)
		//line templates/mail.qtpl:227
		qw422016.E().S(dominio.FiltroSpam)
		//line templates/mail.qtpl:227
		qw422016.N().S(`" )' ><i class="fa fa-list"></i></button>
                                    <button class="btn btn-xs btn-danger" data-toggle="tooltip" data-placement="top" title="Eliminar dominio" onclick="location.href='removeDominio?id=`)
		//line templates/mail.qtpl:228
		qw422016.N().D(int(dominio.ID))
		//line templates/mail.qtpl:228
		qw422016.N().S(`';"><i class="fa fa-trash-o"></i></button>
                                </td>
                            </tr>
                            `)
		//line templates/mail.qtpl:231
	}
	//line templates/mail.qtpl:231
	qw422016.N().S(`
                            </tbody>
                        </table>
                        <a class="btn btn-info btn-sm" href="/" id="webmail">Webmail</a>
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
                    Configuración del dominio de correo electrónico
                </div>
                <div class="panel-body">
                    <form id="adddominio" action="/mail" onsubmit="return validateForm()" role=form method="post">
                        Nombre
                        <input id="id" name="id" id="id" hidden="true" >
                        <input class="form-control" name="nombre" id="nombre">
                        <br>
                        Filtro de correo basura
                        <select class="form-control" id="filtro" name="filtro">
                            <option disabled selected value> -- Elegir tipo de filtro -- </option>
                            <option value="alto">Alto</option>
                            <option value="medio">Medio</option>
                            <option value="bajo">Bajo</option>
                            <option value="off">Desactivado</option>
                        </select>
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
//line templates/mail.qtpl:297
}

//line templates/mail.qtpl:297
func WriteMailTemplate(qq422016 qtio422016.Writer, dominios []model.Dominio, error string) {
	//line templates/mail.qtpl:297
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/mail.qtpl:297
	StreamMailTemplate(qw422016, dominios, error)
	//line templates/mail.qtpl:297
	qt422016.ReleaseWriter(qw422016)
//line templates/mail.qtpl:297
}

//line templates/mail.qtpl:297
func MailTemplate(dominios []model.Dominio, error string) string {
	//line templates/mail.qtpl:297
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/mail.qtpl:297
	WriteMailTemplate(qb422016, dominios, error)
	//line templates/mail.qtpl:297
	qs422016 := string(qb422016.B)
	//line templates/mail.qtpl:297
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/mail.qtpl:297
	return qs422016
//line templates/mail.qtpl:297
}
