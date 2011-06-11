include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=	*.go

include $(GOROOT)/src/Make.pkg

link: package
	$(O)l -o main _go_.$(O)
run: link
	./main
