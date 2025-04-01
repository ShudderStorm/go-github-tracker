package gin

const DefaultPort string = ":8080"

func Run(port string) {
	R.Run(port)
}
