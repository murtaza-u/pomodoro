PREFIX ?= /usr
BINDIR ?= $(PREFIX)/bin

all: pom

pom:
	go build

install: all
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 pom $(DESTDIR)$(BINDIR)

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/pom

clean:
	rm -f pom

.PHONY: all install uninstall clean
