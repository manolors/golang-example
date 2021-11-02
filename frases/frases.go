package frases

import (
	"math/rand"
	"time"
)

// Nueva devuelve aleatoriamente una frase de una película
func Nueva() string {
	frases := []string{
		"Hay tiempo de comer tranquilos...",
		"Que la fuerza te acompañe",
		"Voy a hacerle una oferta que no podrá rechazar",
		"Me encanta el olor del napalm por la mañana",
		"Uno del censo vino a verme. Me comí su hígado con alubias y regado con un delicioso Chianti",
		"Vamos a necesitar un barco más grande",
		"Mamá siempre decía: la vida es como una caja de bombones. Nunca sabes lo que te va a tocar",
		"En ocasiones veo muertos",
		"Quita tus sucias garras de mí… Mono asqueroso",
		"Por favor, Hal, abre las puertas",
		"Son tiempos difíciles para los soñadores",
		"He cruzado océanos de tiempo para encontrarte",
		"Estoy a veinte minutos de allí. Llegaré en diez",
		"He visto cosas que vosotros no creeríais.",
	}
	rand.Seed(time.Now().UnixNano())
	return frases[rand.Intn(len(frases))]
}
