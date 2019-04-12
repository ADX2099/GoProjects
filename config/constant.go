package config

//SYSTEM

const SYSTEM_PORT_DEFAULT = "9010"

// consul service
const CONSUL_REGISTRATION_ID = "laboratorio1-server"
const CONSUL_REGISTRATION_NAME = "laboratorio1-server"
const CONSUL_URL_CHECK = "http://%s:%v/info"
const CONSUL_INTERVAL = "5s"
const CONSUL_TIMEOUT = "3s"

// time manager
const TIMEZONE_DEFAULT = "America/Mexico_City"

//security
const ERROR_CREATE_TOKEN = 1001
const SUSSES_CREATE_TOKEN = 0
const MINUTES_EXPIRE_TOKEN = 1

//process microservices

const OK_PROCESS = 0

// ERRORS

const USER_NOT_AUTHORIZE = 403
const NOT_FOUND = 404
const ERROR_IN_SECURITY_PROCESS = 10001
const ERROR_IN_DATA = 10002
const ERROR_IN_JSON_STRING = 10003

// MESSAGES

const MSG_VERIFY_DATA = "Excepción: Verifique los datos  xxxxx. "
const MSG_ERROR_JSON_STRING = "Excepción: Cadena Json erronea "
const MSG_USER_NOT_AUTHORIZE = "Usuario no autorizado"
