package fan_out

func main() {
	for {
		data := <- src
		select {
		case <- ctx.Done():
		  return ctx.Err()
		case dst1<-data:
		case dst2<-data:
		case dst3<-data:
		}
	}
}
