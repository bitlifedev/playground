package user

import "fmt"

func init() {
	fmt.Println("Init User")
}

func Run() error {
	// instantiate any dependencies my Component struct will need
	// in this example, let's imagine the comment package an implementation
	// that matches what our Service expects
	//commentRepo := Comment.NewRepo(dbConnectionInfo)

	// We can then pass this into our NewService constructor like so:
	//comp, err := NewService(commentRepo)
	//if err != nil {
	// handle this properly with logs/metrics/alerts etc
	//	return err
	//}

	return nil
}
