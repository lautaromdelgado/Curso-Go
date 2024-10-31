package decorator

type Notificador interface {
	Enviar(mensaje string) string
}

type NotificadorBase struct {}

func(n NotificadorBase) Enviar(mensaje string) string {
	return "Mensaje enviado: " + mensaje
}

type NotificadorConAlerta struct{
	Notificador Notificador
}

func(n NotificadorConAlerta) Enviar(mensaje string) string {
	return n.Notificador.Enviar("⚠️ ALERTA: " + mensaje)
}