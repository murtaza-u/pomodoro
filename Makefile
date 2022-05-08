PREFIX ?= /usr
BINDIR ?= $(PREFIX)/bin

all: pomo

pomo:
	go build

install: all
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 pomo $(DESTDIR)$(BINDIR)

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/pomo

clean:
	rm -f pomo

.PHONY: all install uninstall clean
