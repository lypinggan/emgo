void memmove(unsafe$Pointer dst, unsafe$Pointer src, uint n);
void memcpy(unsafe$Pointer dst, unsafe$Pointer src, uint n);
void memset(unsafe$Pointer s, byte b, uint n);

__attribute__ ((noreturn))
void panic(string s);

#define NEW(typ) (typ*) builtin$Alloc(1, sizeof(typ), __alignof__(typ))

#define MAKESLI(typ, l) (slice){							\
	builtin$Alloc(l, sizeof(typ), __alignof__(typ)), l, l	\
}

#define MAKESLIC(typ, l, c) (slice){						\
	builtin$Alloc(c, sizeof(typ), __alignof__(typ)), l, c	\
}