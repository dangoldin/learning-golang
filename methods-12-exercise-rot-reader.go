package main

import (
       "io"
       "os"
       "strings"
)

type rot13Reader struct {
     r io.Reader
}

func rot13(c byte) byte {
     if (c >= 'a' && c <= 'z') {
     	if (c > 'm') {
	      return c - 13
	      	     } else {
		       	    return c + 13	
			    	   }
				   }
				   if (c >= 'A' && c <= 'Z') {
				      if (c > 'M') {
				      	    return c - 13
					    	   } else {
						     	  return c + 13
							  	 }
								 }
								 return 0
}

func (r rot13Reader) Read(a[] byte) (int, error) {
     n, err := r.r.Read(a)
     for i, _ := range a {
     	 a[i] = rot13(a[i])
	 }    
	 return n, err
}

func main() {
     s := strings.NewReader("Lbh penpxrq gur pbqr!")
     r := rot13Reader{s}
     io.Copy(os.Stdout, &r)
}
