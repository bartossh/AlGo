###
CFLAGS  = -std=c99
CFLAGS += -g
CFLAGS += -Wall
CFLAGS += -Wextra
CFLAGS += -pedantic
CFLAGS += -Werror
CFLAGS += -Wmissing-declarations
CFLAGS += -DUNITY_SUPPORT_64 -DUNITY_OUTPUT_COLOR

ASANFLAGS  = -fsanitize=address
ASANFLAGS += -fno-common
ASANFLAGS += -fno-omit-frame-pointer

# Used to test static functions.
UTFLAGS = -DUNIT_TEST


.PHONY: memcheck
memcheck: ./*.c
	@echo Compiling $@
	@$(CC) $(ASANFLAGS) $(CFLAGS) $(UTFLAGS) ./*.c -o solver
	@./solver
	@echo "Memory check passed"

.PHONY: clean
clean:
	rm -rf *.o *.out *.out.dSYM *.pem solver

solve: ./*.c
	@echo Compiling $@
	@$(CC) $(CFLAGS) $(UTFLAGS) ./*.c -o solver
	@./solver

