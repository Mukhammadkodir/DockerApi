package repo

import (
	pb "github/DockerApi/api/genproto/user"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	Get(*pb.ById) (*pb.User, error)
	Delete(*pb.ById) (*pb.Empty, error)
	Update(*pb.User) (*pb.User, error)
}
