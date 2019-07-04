package main

import (
	"fmt"
	"io"
	"strings"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)


func main(){

	// dot, _, err := git.DotGitToOSFilesystems(".", false)
	// if err != nil {
	// 	panic(err)
	// }

	// s := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())
	//ms := memory.NewStorage()

	//f := memfs.New()
	
	repo, err := git.PlainOpen(".")
    // if err != nil {
    //     panic(err)
    // }

	w, err := repo.Worktree()
	// if err != nil {
	// 	panic(err)
	// }
	// 대체 가능 ...
	// idx, err := w.R.Storer.Index()
	// if err != nil {
	// 	panic(err)
	// }
	

	reader := strings.NewReader("ms working !! 9999")
	// 대체 가능 ...
	obj :=  w.R.Storer.NewEncodedObject()
	obj.SetType(plumbing.BlobObject)
	obj.SetSize(reader.Size())

	writer, err := obj.Writer()
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(writer, reader); err != nil {
		panic(err)
	}
	// 대체 가능 ...
	w.R.Storer.SetEncodedObject(obj)
	// h, err :=  w.R.Storer.SetEncodedObject(obj)
	// if err != nil {
	// 	panic(err)
	// }


	//e := idx.Add("helloTest999.readme")	

	//e.Hash = h
	//e.ModifiedAt = info.ModTime()
	//e.Mode, err = filemode.NewFromOSFileMode(info.Mode())
	//e.Mode = 100644

	if err != nil {
		panic(err)
	}

	// 등록 ....  대체 가능 ... 
	//w.R.Storer.SetIndex(idx)

	//fillSystemInfo(e, info.Sys())

	// if e.Mode.IsRegular() {
	// 	e.Size = uint32(info.Size())
	// }



	// commit, err := w.Commit("this is my memory commit6666667", &git.CommitOptions{
	// 	Author: &object.Signature{
	// 		Name:  "hosuk",
	// 		Email: "kirklayer@gmail.com",
	// 		When:  time.Now(),
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// obj2, err := repo.CommitObject(commit)
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println(obj2)
	// io.WriteString(os.Stdout , st.Name())
	fmt.Printf("hello \n")
	//fmt.Println(obj)
}








