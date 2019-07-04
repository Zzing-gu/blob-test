package main

import (
	"fmt"
	"io"
	"strings"
	"time"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)


func main(){


	
	repo, err := git.PlainOpen(".")
    if err != nil {
        panic(err)
    }

	w, err := repo.Worktree()
	if err != nil {
		panic(err)
	}
	
	idx, err := w.R.Storer.Index()
	if err != nil {
		panic(err)
	}
	

	reader := strings.NewReader("ms working !! 9999")
	
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

	
	h, err :=  w.R.Storer.SetEncodedObject(obj)
	if err != nil {
		panic(err)
	}


	e := idx.Add("hello9999.readme")	

	e.Hash = h
	//e.ModifiedAt = info.ModTime()
	//e.Mode, err = filemode.NewFromOSFileMode(info.Mode())
	e.Mode = 100644


	
	w.R.Storer.SetIndex(idx)

	



	commit, err := w.Commit("this is my add commit test", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "hosuk",
			Email: "kirklayer@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		panic(err)
	}

	commitObj, err := repo.CommitObject(commit)
	if err != nil {
		panic(err)
	}
	fmt.Println(commitObj)
	
	fmt.Printf("hello \n")
	
}








