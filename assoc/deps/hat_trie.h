
#include <stdio.h>

void* malloc_or_die(size_t);
void* realloc_or_die(void*, size_t);
FILE* fopen_or_die(const char*, const char*);


#include <stdlib.h>


#include <stddef.h>
#include <limits.h>
#include <signal.h>

/*
 *  For gcc with _STDINT_H, fill in the PRINTF_INT*_MODIFIER macros, and
 *  do nothing else.  On the Mac OS X version of gcc this is _STDINT_H_.
 */

#if ((defined(__STDC__) && __STDC__ && __STDC_VERSION__ >= 199901L) || (defined (__WATCOMC__) && (defined (_STDINT_H_INCLUDED) || __WATCOMC__ >= 1250)) || (defined(__GNUC__) && (defined(_STDINT_H) || defined(_STDINT_H_) || defined (__UINT_FAST64_TYPE__)) )) && !defined (_PSTDINT_H_INCLUDED)
#include <stdint.h>
#define _PSTDINT_H_INCLUDED
# ifndef PRINTF_INT64_MODIFIER
#  define PRINTF_INT64_MODIFIER "ll"
# endif
# ifndef PRINTF_INT32_MODIFIER
#  define PRINTF_INT32_MODIFIER "l"
# endif
# ifndef PRINTF_INT16_MODIFIER
#  define PRINTF_INT16_MODIFIER "h"
# endif
# ifndef PRINTF_INTMAX_MODIFIER
#  define PRINTF_INTMAX_MODIFIER PRINTF_INT64_MODIFIER
# endif
# ifndef PRINTF_INT64_HEX_WIDTH
#  define PRINTF_INT64_HEX_WIDTH "16"
# endif
# ifndef PRINTF_INT32_HEX_WIDTH
#  define PRINTF_INT32_HEX_WIDTH "8"
# endif
# ifndef PRINTF_INT16_HEX_WIDTH
#  define PRINTF_INT16_HEX_WIDTH "4"
# endif
# ifndef PRINTF_INT8_HEX_WIDTH
#  define PRINTF_INT8_HEX_WIDTH "2"
# endif
# ifndef PRINTF_INT64_DEC_WIDTH
#  define PRINTF_INT64_DEC_WIDTH "20"
# endif
# ifndef PRINTF_INT32_DEC_WIDTH
#  define PRINTF_INT32_DEC_WIDTH "10"
# endif
# ifndef PRINTF_INT16_DEC_WIDTH
#  define PRINTF_INT16_DEC_WIDTH "5"
# endif
# ifndef PRINTF_INT8_DEC_WIDTH
#  define PRINTF_INT8_DEC_WIDTH "3"
# endif
# ifndef PRINTF_INTMAX_HEX_WIDTH
#  define PRINTF_INTMAX_HEX_WIDTH PRINTF_INT64_HEX_WIDTH
# endif
# ifndef PRINTF_INTMAX_DEC_WIDTH
#  define PRINTF_INTMAX_DEC_WIDTH PRINTF_INT64_DEC_WIDTH
# endif

/*
 *  Something really weird is going on with Open Watcom.  Just pull some of
 *  these duplicated definitions from Open Watcom's stdint.h file for now.
 */

# if defined (__WATCOMC__) && __WATCOMC__ >= 1250
#  if !defined (INT64_C)
#   define INT64_C(x)   (x + (INT64_MAX - INT64_MAX))
#  endif
#  if !defined (UINT64_C)
#   define UINT64_C(x)  (x + (UINT64_MAX - UINT64_MAX))
#  endif
#  if !defined (INT32_C)
#   define INT32_C(x)   (x + (INT32_MAX - INT32_MAX))
#  endif
#  if !defined (UINT32_C)
#   define UINT32_C(x)  (x + (UINT32_MAX - UINT32_MAX))
#  endif
#  if !defined (INT16_C)
#   define INT16_C(x)   (x)
#  endif
#  if !defined (UINT16_C)
#   define UINT16_C(x)  (x)
#  endif
#  if !defined (INT8_C)
#   define INT8_C(x)   (x)
#  endif
#  if !defined (UINT8_C)
#   define UINT8_C(x)  (x)
#  endif
#  if !defined (UINT64_MAX)
#   define UINT64_MAX  18446744073709551615ULL
#  endif
#  if !defined (INT64_MAX)
#   define INT64_MAX  9223372036854775807LL
#  endif
#  if !defined (UINT32_MAX)
#   define UINT32_MAX  4294967295UL
#  endif
#  if !defined (INT32_MAX)
#   define INT32_MAX  2147483647L
#  endif
#  if !defined (INTMAX_MAX)
#   define INTMAX_MAX INT64_MAX
#  endif
#  if !defined (INTMAX_MIN)
#   define INTMAX_MIN INT64_MIN
#  endif
# endif
#endif

#ifndef _PSTDINT_H_INCLUDED
#define _PSTDINT_H_INCLUDED

#ifndef SIZE_MAX
# define SIZE_MAX (~(size_t)0)
#endif

/*
 *  Deduce the type assignments from limits.h under the assumption that
 *  integer sizes in bits are powers of 2, and follow the ANSI
 *  definitions.
 */

#ifndef UINT8_MAX
# define UINT8_MAX 0xff
#endif
#ifndef uint8_t
# if (UCHAR_MAX == UINT8_MAX) || defined (S_SPLINT_S)
    typedef unsigned char uint8_t;
#   define UINT8_C(v) ((uint8_t) v)
# else
#   error "Platform not supported"
# endif
#endif

#ifndef INT8_MAX
# define INT8_MAX 0x7f
#endif
#ifndef INT8_MIN
# define INT8_MIN INT8_C(0x80)
#endif
#ifndef int8_t
# if (SCHAR_MAX == INT8_MAX) || defined (S_SPLINT_S)
    typedef signed char int8_t;
#   define INT8_C(v) ((int8_t) v)
# else
#   error "Platform not supported"
# endif
#endif

#ifndef UINT16_MAX
# define UINT16_MAX 0xffff
#endif
#ifndef uint16_t
#if (UINT_MAX == UINT16_MAX) || defined (S_SPLINT_S)
  typedef unsigned int uint16_t;
# ifndef PRINTF_INT16_MODIFIER
#  define PRINTF_INT16_MODIFIER ""
# endif
# define UINT16_C(v) ((uint16_t) (v))
#elif (USHRT_MAX == UINT16_MAX)
  typedef unsigned short uint16_t;
# define UINT16_C(v) ((uint16_t) (v))
# ifndef PRINTF_INT16_MODIFIER
#  define PRINTF_INT16_MODIFIER "h"
# endif
#else
#error "Platform not supported"
#endif
#endif

#ifndef INT16_MAX
# define INT16_MAX 0x7fff
#endif
#ifndef INT16_MIN
# define INT16_MIN INT16_C(0x8000)
#endif
#ifndef int16_t
#if (INT_MAX == INT16_MAX) || defined (S_SPLINT_S)
  typedef signed int int16_t;
# define INT16_C(v) ((int16_t) (v))
# ifndef PRINTF_INT16_MODIFIER
#  define PRINTF_INT16_MODIFIER ""
# endif
#elif (SHRT_MAX == INT16_MAX)
  typedef signed short int16_t;
# define INT16_C(v) ((int16_t) (v))
# ifndef PRINTF_INT16_MODIFIER
#  define PRINTF_INT16_MODIFIER "h"
# endif
#else
#error "Platform not supported"
#endif
#endif

#ifndef UINT32_MAX
# define UINT32_MAX (0xffffffffUL)
#endif
#ifndef uint32_t
#if (ULONG_MAX == UINT32_MAX) || defined (S_SPLINT_S)
  typedef unsigned long uint32_t;
# define UINT32_C(v) v ## UL
# ifndef PRINTF_INT32_MODIFIER
#  define PRINTF_INT32_MODIFIER "l"
# endif
#elif (UINT_MAX == UINT32_MAX)
  typedef unsigned int uint32_t;
# ifndef PRINTF_INT32_MODIFIER
#  define PRINTF_INT32_MODIFIER ""
# endif
# define UINT32_C(v) v ## U
#elif (USHRT_MAX == UINT32_MAX)
  typedef unsigned short uint32_t;
# define UINT32_C(v) ((unsigned short) (v))
# ifndef PRINTF_INT32_MODIFIER
#  define PRINTF_INT32_MODIFIER ""
# endif
#else
#error "Platform not supported"
#endif
#endif

#ifndef INT32_MAX
# define INT32_MAX (0x7fffffffL)
#endif
#ifndef INT32_MIN
# define INT32_MIN INT32_C(0x80000000)
#endif
#ifndef int32_t
#if (LONG_MAX == INT32_MAX) || defined (S_SPLINT_S)
  typedef signed long int32_t;
# define INT32_C(v) v ## L
# ifndef PRINTF_INT32_MODIFIER
#  define PRINTF_INT32_MODIFIER "l"
# endif
#elif (INT_MAX == INT32_MAX)
  typedef signed int int32_t;
# define INT32_C(v) v
# ifndef PRINTF_INT32_MODIFIER
#  define PRINTF_INT32_MODIFIER ""
# endif
#elif (SHRT_MAX == INT32_MAX)
  typedef signed short int32_t;
# define INT32_C(v) ((short) (v))
# ifndef PRINTF_INT32_MODIFIER
#  define PRINTF_INT32_MODIFIER ""
# endif
#else
#error "Platform not supported"
#endif
#endif

/*
 *  The macro stdint_int64_defined is temporarily used to record
 *  whether or not 64 integer support is available.  It must be
 *  defined for any 64 integer extensions for new platforms that are
 *  added.
 */

#undef stdint_int64_defined
#if (defined(__STDC__) && defined(__STDC_VERSION__)) || defined (S_SPLINT_S)
# if (__STDC__ && __STDC_VERSION__ >= 199901L) || defined (S_SPLINT_S)
#  define stdint_int64_defined
   typedef long long int64_t;
   typedef unsigned long long uint64_t;
#  define UINT64_C(v) v ## ULL
#  define  INT64_C(v) v ## LL
#  ifndef PRINTF_INT64_MODIFIER
#   define PRINTF_INT64_MODIFIER "ll"
#  endif
# endif
#endif

#if !defined (stdint_int64_defined)
# if defined(__GNUC__)
#  define stdint_int64_defined
   __extension__ typedef long long int64_t;
   __extension__ typedef unsigned long long uint64_t;
#  define UINT64_C(v) v ## ULL
#  define  INT64_C(v) v ## LL
#  ifndef PRINTF_INT64_MODIFIER
#   define PRINTF_INT64_MODIFIER "ll"
#  endif
# elif defined(__MWERKS__) || defined (__SUNPRO_C) || defined (__SUNPRO_CC) || defined (__APPLE_CC__) || defined (_LONG_LONG) || defined (_CRAYC) || defined (S_SPLINT_S)
#  define stdint_int64_defined
   typedef long long int64_t;
   typedef unsigned long long uint64_t;
#  define UINT64_C(v) v ## ULL
#  define  INT64_C(v) v ## LL
#  ifndef PRINTF_INT64_MODIFIER
#   define PRINTF_INT64_MODIFIER "ll"
#  endif
# elif (defined(__WATCOMC__) && defined(__WATCOM_INT64__)) || (defined(_MSC_VER) && _INTEGRAL_MAX_BITS >= 64) || (defined (__BORLANDC__) && __BORLANDC__ > 0x460) || defined (__alpha) || defined (__DECC)
#  define stdint_int64_defined
   typedef __int64 int64_t;
   typedef unsigned __int64 uint64_t;
#  define UINT64_C(v) v ## UI64
#  define  INT64_C(v) v ## I64
#  ifndef PRINTF_INT64_MODIFIER
#   define PRINTF_INT64_MODIFIER "I64"
#  endif
# endif
#endif

#if !defined (LONG_LONG_MAX) && defined (INT64_C)
# define LONG_LONG_MAX INT64_C (9223372036854775807)
#endif
#ifndef ULONG_LONG_MAX
# define ULONG_LONG_MAX UINT64_C (18446744073709551615)
#endif

#if !defined (INT64_MAX) && defined (INT64_C)
# define INT64_MAX INT64_C (9223372036854775807)
#endif
#if !defined (INT64_MIN) && defined (INT64_C)
# define INT64_MIN INT64_C (-9223372036854775808)
#endif
#if !defined (UINT64_MAX) && defined (INT64_C)
# define UINT64_MAX UINT64_C (18446744073709551615)
#endif

/*
 *  Width of hexadecimal for number field.
 */

#ifndef PRINTF_INT64_HEX_WIDTH
# define PRINTF_INT64_HEX_WIDTH "16"
#endif
#ifndef PRINTF_INT32_HEX_WIDTH
# define PRINTF_INT32_HEX_WIDTH "8"
#endif
#ifndef PRINTF_INT16_HEX_WIDTH
# define PRINTF_INT16_HEX_WIDTH "4"
#endif
#ifndef PRINTF_INT8_HEX_WIDTH
# define PRINTF_INT8_HEX_WIDTH "2"
#endif

#ifndef PRINTF_INT64_DEC_WIDTH
# define PRINTF_INT64_DEC_WIDTH "20"
#endif
#ifndef PRINTF_INT32_DEC_WIDTH
# define PRINTF_INT32_DEC_WIDTH "10"
#endif
#ifndef PRINTF_INT16_DEC_WIDTH
# define PRINTF_INT16_DEC_WIDTH "5"
#endif
#ifndef PRINTF_INT8_DEC_WIDTH
# define PRINTF_INT8_DEC_WIDTH "3"
#endif

/*
 *  Ok, lets not worry about 128 bit integers for now.  Moore's law says
 *  we don't need to worry about that until about 2040 at which point
 *  we'll have bigger things to worry about.
 */

#ifdef stdint_int64_defined
  typedef int64_t intmax_t;
  typedef uint64_t uintmax_t;
# define  INTMAX_MAX   INT64_MAX
# define  INTMAX_MIN   INT64_MIN
# define UINTMAX_MAX  UINT64_MAX
# define UINTMAX_C(v) UINT64_C(v)
# define  INTMAX_C(v)  INT64_C(v)
# ifndef PRINTF_INTMAX_MODIFIER
#   define PRINTF_INTMAX_MODIFIER PRINTF_INT64_MODIFIER
# endif
# ifndef PRINTF_INTMAX_HEX_WIDTH
#  define PRINTF_INTMAX_HEX_WIDTH PRINTF_INT64_HEX_WIDTH
# endif
# ifndef PRINTF_INTMAX_DEC_WIDTH
#  define PRINTF_INTMAX_DEC_WIDTH PRINTF_INT64_DEC_WIDTH
# endif
#else
  typedef int32_t intmax_t;
  typedef uint32_t uintmax_t;
# define  INTMAX_MAX   INT32_MAX
# define UINTMAX_MAX  UINT32_MAX
# define UINTMAX_C(v) UINT32_C(v)
# define  INTMAX_C(v)  INT32_C(v)
# ifndef PRINTF_INTMAX_MODIFIER
#   define PRINTF_INTMAX_MODIFIER PRINTF_INT32_MODIFIER
# endif
# ifndef PRINTF_INTMAX_HEX_WIDTH
#  define PRINTF_INTMAX_HEX_WIDTH PRINTF_INT32_HEX_WIDTH
# endif
# ifndef PRINTF_INTMAX_DEC_WIDTH
#  define PRINTF_INTMAX_DEC_WIDTH PRINTF_INT32_DEC_WIDTH
# endif
#endif

/*
 *  Because this file currently only supports platforms which have
 *  precise powers of 2 as bit sizes for the default integers, the
 *  least definitions are all trivial.  Its possible that a future
 *  version of this file could have different definitions.
 */

#ifndef stdint_least_defined
  typedef   int8_t   int_least8_t;
  typedef  uint8_t  uint_least8_t;
  typedef  int16_t  int_least16_t;
  typedef uint16_t uint_least16_t;
  typedef  int32_t  int_least32_t;
  typedef uint32_t uint_least32_t;
# define PRINTF_LEAST32_MODIFIER PRINTF_INT32_MODIFIER
# define PRINTF_LEAST16_MODIFIER PRINTF_INT16_MODIFIER
# define  UINT_LEAST8_MAX  UINT8_MAX
# define   INT_LEAST8_MAX   INT8_MAX
# define UINT_LEAST16_MAX UINT16_MAX
# define  INT_LEAST16_MAX  INT16_MAX
# define UINT_LEAST32_MAX UINT32_MAX
# define  INT_LEAST32_MAX  INT32_MAX
# define   INT_LEAST8_MIN   INT8_MIN
# define  INT_LEAST16_MIN  INT16_MIN
# define  INT_LEAST32_MIN  INT32_MIN
# ifdef stdint_int64_defined
    typedef  int64_t  int_least64_t;
    typedef uint64_t uint_least64_t;
#   define PRINTF_LEAST64_MODIFIER PRINTF_INT64_MODIFIER
#   define UINT_LEAST64_MAX UINT64_MAX
#   define  INT_LEAST64_MAX  INT64_MAX
#   define  INT_LEAST64_MIN  INT64_MIN
# endif
#endif
#undef stdint_least_defined

/*
 *  The ANSI C committee pretending to know or specify anything about
 *  performance is the epitome of misguided arrogance.  The mandate of
 *  this file is to *ONLY* ever support that absolute minimum
 *  definition of the fast integer types, for compatibility purposes.
 *  No extensions, and no attempt to suggest what may or may not be a
 *  faster integer type will ever be made in this file.  Developers are
 *  warned to stay away from these types when using this or any other
 *  stdint.h.
 */

typedef   int_least8_t   int_fast8_t;
typedef  uint_least8_t  uint_fast8_t;
typedef  int_least16_t  int_fast16_t;
typedef uint_least16_t uint_fast16_t;
typedef  int_least32_t  int_fast32_t;
typedef uint_least32_t uint_fast32_t;
#define  UINT_FAST8_MAX  UINT_LEAST8_MAX
#define   INT_FAST8_MAX   INT_LEAST8_MAX
#define UINT_FAST16_MAX UINT_LEAST16_MAX
#define  INT_FAST16_MAX  INT_LEAST16_MAX
#define UINT_FAST32_MAX UINT_LEAST32_MAX
#define  INT_FAST32_MAX  INT_LEAST32_MAX
#define   INT_FAST8_MIN   INT_LEAST8_MIN
#define  INT_FAST16_MIN  INT_LEAST16_MIN
#define  INT_FAST32_MIN  INT_LEAST32_MIN
#ifdef stdint_int64_defined
  typedef  int_least64_t  int_fast64_t;
  typedef uint_least64_t uint_fast64_t;
# define UINT_FAST64_MAX UINT_LEAST64_MAX
# define  INT_FAST64_MAX  INT_LEAST64_MAX
# define  INT_FAST64_MIN  INT_LEAST64_MIN
#endif

#undef stdint_int64_defined

/*
 *  Whatever piecemeal, per compiler thing we can do about the wchar_t
 *  type limits.
 */

#if defined(__WATCOMC__) || defined(_MSC_VER) || defined (__GNUC__)
# include <wchar.h>
# ifndef WCHAR_MIN
#  define WCHAR_MIN 0
# endif
# ifndef WCHAR_MAX
#  define WCHAR_MAX ((wchar_t)-1)
# endif
#endif

/*
 *  Whatever piecemeal, per compiler/platform thing we can do about the
 *  (u)intptr_t types and limits.
 */

#if defined (_MSC_VER) && defined (_UINTPTR_T_DEFINED)
# define STDINT_H_UINTPTR_T_DEFINED
#endif

#ifndef STDINT_H_UINTPTR_T_DEFINED
# if defined (__alpha__) || defined (__ia64__) || defined (__x86_64__) || defined (_WIN64)
#  define stdint_intptr_bits 64
# elif defined (__WATCOMC__) || defined (__TURBOC__)
#  if defined(__TINY__) || defined(__SMALL__) || defined(__MEDIUM__)
#    define stdint_intptr_bits 16
#  else
#    define stdint_intptr_bits 32
#  endif
# elif defined (__i386__) || defined (_WIN32) || defined (WIN32)
#  define stdint_intptr_bits 32
# elif defined (__INTEL_COMPILER)
/* TODO -- what did Intel do about x86-64? */
# endif

# ifdef stdint_intptr_bits
#  define stdint_intptr_glue3_i(a,b,c)  a##b##c
#  define stdint_intptr_glue3(a,b,c)    stdint_intptr_glue3_i(a,b,c)
#  ifndef PRINTF_INTPTR_MODIFIER
#    define PRINTF_INTPTR_MODIFIER      stdint_intptr_glue3(PRINTF_INT,stdint_intptr_bits,_MODIFIER)
#  endif
#  ifndef PTRDIFF_MAX
#    define PTRDIFF_MAX                 stdint_intptr_glue3(INT,stdint_intptr_bits,_MAX)
#  endif
#  ifndef PTRDIFF_MIN
#    define PTRDIFF_MIN                 stdint_intptr_glue3(INT,stdint_intptr_bits,_MIN)
#  endif
#  ifndef UINTPTR_MAX
#    define UINTPTR_MAX                 stdint_intptr_glue3(UINT,stdint_intptr_bits,_MAX)
#  endif
#  ifndef INTPTR_MAX
#    define INTPTR_MAX                  stdint_intptr_glue3(INT,stdint_intptr_bits,_MAX)
#  endif
#  ifndef INTPTR_MIN
#    define INTPTR_MIN                  stdint_intptr_glue3(INT,stdint_intptr_bits,_MIN)
#  endif
#  ifndef INTPTR_C
#    define INTPTR_C(x)                 stdint_intptr_glue3(INT,stdint_intptr_bits,_C)(x)
#  endif
#  ifndef UINTPTR_C
#    define UINTPTR_C(x)                stdint_intptr_glue3(UINT,stdint_intptr_bits,_C)(x)
#  endif
  typedef stdint_intptr_glue3(uint,stdint_intptr_bits,_t) uintptr_t;
  typedef stdint_intptr_glue3( int,stdint_intptr_bits,_t)  intptr_t;
# else
/* TODO -- This following is likely wrong for some platforms, and does
   nothing for the definition of uintptr_t. */
  typedef ptrdiff_t intptr_t;
# endif
# define STDINT_H_UINTPTR_T_DEFINED
#endif

/*
 *  Assumes sig_atomic_t is signed and we have a 2s complement machine.
 */

#ifndef SIG_ATOMIC_MAX
# define SIG_ATOMIC_MAX ((((sig_atomic_t) 1) << (sizeof (sig_atomic_t)*CHAR_BIT-1)) - 1)
#endif

#endif

#if defined (__TEST_PSTDINT_FOR_CORRECTNESS)

/* 
 *  Please compile with the maximum warning settings to make sure macros are not
 *  defined more than once.
 */
 
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
 
#define glue3_aux(x,y,z) x ## y ## z
#define glue3(x,y,z) glue3_aux(x,y,z)

#define DECLU(bits) glue3(uint,bits,_t) glue3(u,bits,=) glue3(UINT,bits,_C) (0);
#define DECLI(bits) glue3(int,bits,_t) glue3(i,bits,=) glue3(INT,bits,_C) (0);

#define DECL(us,bits) glue3(DECL,us,) (bits)

#define TESTUMAX(bits) glue3(u,bits,=) glue3(~,u,bits); if (glue3(UINT,bits,_MAX) glue3(!=,u,bits)) printf ("Something wrong with UINT%d_MAX\n", bits)
 
int main () {
	DECL(I,8)
	DECL(U,8)
	DECL(I,16)
	DECL(U,16)
	DECL(I,32)
	DECL(U,32)
#ifdef INT64_MAX
	DECL(I,64)
	DECL(U,64)
#endif
	intmax_t imax = INTMAX_C(0);
	uintmax_t umax = UINTMAX_C(0);
	char str0[256], str1[256];

	sprintf (str0, "%d %x\n", 0, ~0);
	
	sprintf (str1, "%d %x\n",  i8, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with i8 : %s\n", str1);
	sprintf (str1, "%u %x\n",  u8, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with u8 : %s\n", str1);
	sprintf (str1, "%d %x\n",  i16, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with i16 : %s\n", str1);
	sprintf (str1, "%u %x\n",  u16, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with u16 : %s\n", str1);	
	sprintf (str1, "%" PRINTF_INT32_MODIFIER "d %x\n",  i32, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with i32 : %s\n", str1);
	sprintf (str1, "%" PRINTF_INT32_MODIFIER "u %x\n",  u32, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with u32 : %s\n", str1);
#ifdef INT64_MAX	
	sprintf (str1, "%" PRINTF_INT64_MODIFIER "d %x\n",  i64, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with i64 : %s\n", str1);
#endif
	sprintf (str1, "%" PRINTF_INTMAX_MODIFIER "d %x\n",  imax, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with imax : %s\n", str1);
	sprintf (str1, "%" PRINTF_INTMAX_MODIFIER "u %x\n",  umax, ~0);
	if (0 != strcmp (str0, str1)) printf ("Something wrong with umax : %s\n", str1);	
	
	TESTUMAX(8);
	TESTUMAX(16);
	TESTUMAX(32);
#ifdef INT64_MAX
	TESTUMAX(64);
#endif

	return EXIT_SUCCESS;
}

#endif


/*
 * This file is part of hat-trie
 *
 * Copyright (c) 2011 by Daniel C. Jones <dcjones@cs.washington.edu>
 *
 *
 * This is an implementation of the HAT-trie data structure described in,
 *
 *    Askitis, N., & Sinha, R. (2007). HAT-trie: a cache-conscious trie-based data
 *    structure for strings. Proceedings of the thirtieth Australasian conference on
 *    Computer science-Volume 62 (pp. 97–105). Australian Computer Society, Inc.
 *
 * The HAT-trie is in essence a hybrid data structure, combining tries and hash
 * tables in a clever way to try to get the best of both worlds.
 *
 */

#ifndef HATTRIE_HATTRIE_H
#define HATTRIE_HATTRIE_H

#ifdef __cplusplus
extern "C" {
#endif

typedef unsigned long value_t;

#include <stdlib.h>
#include <stdbool.h>

typedef struct hattrie_t_ hattrie_t;

hattrie_t* hattrie_create (void);             //< Create an empty hat-trie.
void       hattrie_free   (hattrie_t*);       //< Free all memory used by a trie.
hattrie_t* hattrie_dup    (const hattrie_t*); //< Duplicate an existing trie.
void       hattrie_clear  (hattrie_t*);       //< Remove all entries.

/** number of inserted keys
 */
size_t hattrie_size (hattrie_t*);

/** Find the given key in the trie, inserting it if it does not exist, and
 * returning a pointer to it's key.
 *
 * This pointer is not guaranteed to be valid after additional calls to
 * hattrie_get, hattrie_del, hattrie_clear, or other functions that modifies the
 * trie.
 */
value_t* hattrie_get (hattrie_t*, const char* key, size_t len);

/** Find a given key in the table, returning a NULL pointer if it does not
 * exist. */
value_t* hattrie_tryget (hattrie_t*, const char* key, size_t len);

/** hattrie_walk callback signature */
typedef int (*hattrie_walk_cb)(const char* key, size_t len, value_t* val, void* user_data);

/** hattrie_walk callback return values, controls whether should stop the walk or not */
#define hattrie_walk_stop 0
#define hattrie_walk_continue 1

/** Find stored keys which are prefices of key, and invoke callback for every found key and val.
 *  The invocation order is: short key to long key.
 */
void hattrie_walk (hattrie_t*, const char* key, size_t len, void* user_data, hattrie_walk_cb);

/** Delete a given key from trie. Returns 0 if successful or -1 if not found.
 */
int hattrie_del(hattrie_t* T, const char* key, size_t len);

typedef struct hattrie_iter_t_ hattrie_iter_t;

hattrie_iter_t* hattrie_iter_begin     (const hattrie_t*, bool sorted);
void            hattrie_iter_next      (hattrie_iter_t*);
bool            hattrie_iter_finished  (hattrie_iter_t*);
void            hattrie_iter_free      (hattrie_iter_t*);
const char*     hattrie_iter_key       (hattrie_iter_t*, size_t* len);
value_t*        hattrie_iter_val       (hattrie_iter_t*);

/** Note the hattrie_iter_key() for prefixed search gets the suffix instead of the whole key
 */
hattrie_iter_t* hattrie_iter_with_prefix(const hattrie_t*, bool sorted, const char* prefix, size_t prefix_len);

#ifdef __cplusplus
}
#endif

#endif

/*
 * This file is part of hat-trie.
 *
 * Copyright (c) 2011 by Daniel C. Jones <dcjones@cs.washington.edu>
 *
 */


/*
 * This file is part of hat-trie.
 *
 * Copyright (c) 2011 by Daniel C. Jones <dcjones@cs.washington.edu>
 *
 *
 * This is an implementation of the 'cache-conscious' hash tables described in,
 *
 *    Askitis, N., & Zobel, J. (2005). Cache-conscious collision resolution in
 *    string hash tables. String Processing and Information Retrieval (pp.
 *    91–102). Springer.
 *
 * Briefly, the idea is, as opposed to separate chaining with linked lists, to
 * store keys contiguously in one big array, thereby improving the caching
 * behavior, and reducing space requirments.
 *
 */

#ifndef HATTRIE_AHTABLE_H
#define HATTRIE_AHTABLE_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>
#include <stdbool.h>

typedef unsigned char* slot_t;

typedef struct ahtable_t_
{
    /* these fields are reserved for hattrie to fiddle with */
    uint8_t flag; 
    unsigned char c0;
    unsigned char c1;

    size_t n;        // number of slots
    size_t m;        // numbur of key/value pairs stored
    size_t max_m;    // number of stored keys before we resize

    size_t*  slot_sizes;
    slot_t*  slots;
} ahtable_t;

extern const double ahtable_max_load_factor;
extern const size_t ahtable_initial_size;

ahtable_t* ahtable_create   (void);         // Create an empty hash table.
ahtable_t* ahtable_create_n (size_t n);     // Create an empty hash table, with
                                            //  n slots reserved.

void       ahtable_free   (ahtable_t*);       // Free all memory used by a table.
void       ahtable_clear  (ahtable_t*);       // Remove all entries.
size_t     ahtable_size   (const ahtable_t*); // Number of stored keys.


/** Find the given key in the table, inserting it if it does not exist, and
 * returning a pointer to it's key.
 *
 * This pointer is not guaranteed to be valid after additional calls to
 * ahtable_get, ahtable_del, ahtable_clear, or other functions that modifies the
 * table.
 */
value_t* ahtable_get (ahtable_t*, const char* key, size_t len);


/** Find a given key in the table, returning a NULL pointer if it does not
 * exist. */
value_t* ahtable_tryget (ahtable_t*, const char* key, size_t len);


int ahtable_del(ahtable_t*, const char* key, size_t len);


typedef struct ahtable_iter_t_ ahtable_iter_t;

ahtable_iter_t* ahtable_iter_begin     (const ahtable_t*, bool sorted);
void            ahtable_iter_next      (ahtable_iter_t*);
bool            ahtable_iter_finished  (ahtable_iter_t*);
void            ahtable_iter_free      (ahtable_iter_t*);
const char*     ahtable_iter_key       (ahtable_iter_t*, size_t* len);
value_t*        ahtable_iter_val       (ahtable_iter_t*);


#ifdef __cplusplus
}
#endif

#endif

/*
 * This file is part of hat-trie.
 *
 * Copyright (c) 2011 by Daniel C. Jones <dcjones@cs.washington.edu>
 *
 */



static inline uint32_t fmix(uint32_t h)
{
    h ^= h >> 16;
    h *= 0x85ebca6b;
    h ^= h >> 13;
    h *= 0xc2b2ae35;
    h ^= h >> 16;

    return h;
}


static inline uint32_t rotl32(uint32_t x, int8_t r)
{
    return (x << r) | (x >> (32 - r));
}


uint32_t hash(const char* data, size_t len_)
{
    const int len = (int) len_;
    const int nblocks = len / 4;

    uint32_t h1 = 0xc062fb4a;

    uint32_t c1 = 0xcc9e2d51;
    uint32_t c2 = 0x1b873593;

    //----------
    // body

    const uint32_t * blocks = (const uint32_t*) (data + nblocks * 4);

    int i;
    for(i = -nblocks; i; i++)
    {
        uint32_t k1 = blocks[i];

        k1 *= c1;
        k1 = rotl32(k1, 15);
        k1 *= c2;

        h1 ^= k1;
        h1 = rotl32(h1, 13);
        h1 = h1*5+0xe6546b64;
    }

    //----------
    // tail

    const uint8_t * tail = (const uint8_t*)(data + nblocks*4);

    uint32_t k1 = 0;

    switch(len & 3)
    {
        case 3: k1 ^= tail[2] << 16;
        case 2: k1 ^= tail[1] << 8;
        case 1: k1 ^= tail[0];
              k1 *= c1; k1 = rotl32(k1,15); k1 *= c2; h1 ^= k1;
    }

    //----------
    // finalization

    h1 ^= len;

    h1 = fmix(h1);

    return h1;
}




#include <assert.h>
#include <string.h>



const double ahtable_max_load_factor = 100000.0; /* arbitrary large number => don't resize */
const size_t ahtable_initial_size = 4096;

static size_t keylen(slot_t s) {
    if (0x1 & *s) {
        return (size_t) (*((uint16_t*) s) >> 1);
    }
    else {
        return (size_t) (*s >> 1);
    }
}


ahtable_t* ahtable_create()
{
    return ahtable_create_n(ahtable_initial_size);
}


ahtable_t* ahtable_create_n(size_t n)
{
    ahtable_t* T = (ahtable_t *) malloc_or_die(sizeof(ahtable_t));
    T->flag = 0;
    T->c0 = T->c1 = '\0';

    T->n = n;
    T->m = 0;
    T->max_m = (size_t) (ahtable_max_load_factor * (double) T->n);
    T->slots = (slot_t *) malloc_or_die(n * sizeof(slot_t));
    memset(T->slots, 0, n * sizeof(slot_t));

    T->slot_sizes = (size_t *) malloc_or_die(n * sizeof(size_t));
    memset(T->slot_sizes, 0, n * sizeof(size_t));

    return T;
}


void ahtable_free(ahtable_t* T)
{
    if (T == NULL) return;
    size_t i;
    for (i = 0; i < T->n; ++i) free(T->slots[i]);
    free(T->slots);
    free(T->slot_sizes);
    free(T);
}


size_t ahtable_size(const ahtable_t* T)
{
    return T->m;
}


void ahtable_clear(ahtable_t* T)
{
    size_t i;
    for (i = 0; i < T->n; ++i) free(T->slots[i]);
    T->n = ahtable_initial_size;
    T->slots = (slot_t *) realloc_or_die(T->slots, T->n * sizeof(slot_t));
    memset(T->slots, 0, T->n * sizeof(slot_t));

    T->slot_sizes = (size_t *) realloc_or_die(T->slot_sizes, T->n * sizeof(size_t));
    memset(T->slot_sizes, 0, T->n * sizeof(size_t));
}


static slot_t ins_key(slot_t s, const char* key, size_t len, value_t** val)
{
    // key length
    if (len < 128) {
        s[0] = (unsigned char) (len << 1);
        s += 1;
    }
    else {
        /* The most significant bit is set to indicate that two bytes are
         * being used to store the key length. */
        *((uint16_t*) s) = ((uint16_t) len << 1) | 0x1;
        s += 2;
    }

    // key
    memcpy(s, key, len * sizeof(unsigned char));
    s += len;

    // value
    *val = (value_t*) s;
    **val = 0;
    s += sizeof(value_t);

    return s;
}


static void ahtable_expand(ahtable_t* T)
{
    /* Resizing a table is essentially building a brand new one.
     * One little shortcut we can take on the memory allocation front is to
     * figure out how much memory each slot needs in advance.
     */
    assert(T->n > 0);
    size_t new_n = 2 * T->n;
    size_t* slot_sizes = (size_t *) malloc_or_die(new_n * sizeof(size_t));
    memset(slot_sizes, 0, new_n * sizeof(size_t));

    const char* key;
    size_t len = 0;
    size_t m = 0;
    ahtable_iter_t* i = ahtable_iter_begin(T, false);
    while (!ahtable_iter_finished(i)) {
        key = ahtable_iter_key(i, &len);
        slot_sizes[hash(key, len) % new_n] +=
            len + sizeof(value_t) + (len >= 128 ? 2 : 1);

        ++m;
        ahtable_iter_next(i);
    }
    assert(m == T->m);
    ahtable_iter_free(i);


    /* allocate slots */
    slot_t* slots = (slot_t *) malloc_or_die(new_n * sizeof(slot_t));
    size_t j;
    for (j = 0; j < new_n; ++j) {
        if (slot_sizes[j] > 0) {
            slots[j] = (slot_t) malloc_or_die(slot_sizes[j]);
        }
        else slots[j] = NULL;
    }

    /* rehash values. A few shortcuts can be taken here as well, as we know
     * there will be no collisions. Instead of the regular insertion routine,
     * we keep track of the ends of every slot and simply insert keys.
     * */
    slot_t* slots_next = (slot_t *) malloc_or_die(new_n * sizeof(slot_t));
    memcpy(slots_next, slots, new_n * sizeof(slot_t));
    size_t h;
    m = 0;
    value_t* u;
    value_t* v;
    i = ahtable_iter_begin(T, false);
    while (!ahtable_iter_finished(i)) {

        key = ahtable_iter_key(i, &len);
        h = hash(key, len) % new_n;

        slots_next[h] = ins_key(slots_next[h], key, len, &u);
        v = ahtable_iter_val(i);
        *u = *v;

        ++m;
        ahtable_iter_next(i);
    }
    assert(m == T->m);
    ahtable_iter_free(i);


    free(slots_next);
    for (j = 0; j < T->n; ++j) free(T->slots[j]);

    free(T->slots);
    T->slots = slots;

    free(T->slot_sizes);
    T->slot_sizes = slot_sizes;

    T->n = new_n;
    T->max_m = (size_t) (ahtable_max_load_factor * (double) T->n);
}


static value_t* get_key(ahtable_t* T, const char* key, size_t len, bool insert_missing)
{
    /* if we are at capacity, preemptively resize */
    if (insert_missing && T->m >= T->max_m) {
        ahtable_expand(T);
    }


    uint32_t i = hash(key, len) % T->n;
    size_t k;
    slot_t s;
    value_t* val;

    /* search the array for our key */
    s = T->slots[i];
    while ((size_t) (s - T->slots[i]) < T->slot_sizes[i]) {
        /* get the key length */
        k = keylen(s);
        s += k < 128 ? 1 : 2;

        /* skip keys that are longer than ours */
        if (k != len) {
            s += k + sizeof(value_t);
            continue;
        }

        /* key found. */
        if (memcmp(s, key, len) == 0) {
            return (value_t*) (s + len);
        }
        /* key not found. */
        else {
            s += k + sizeof(value_t);
            continue;
        }
    }


    if (insert_missing) {
        /* the key was not found, so we must insert it. */
        size_t new_size = T->slot_sizes[i];
        new_size += 1 + (len >= 128 ? 1 : 0);    // key length
        new_size += len * sizeof(unsigned char); // key
        new_size += sizeof(value_t);             // value

        T->slots[i] = (slot_t) realloc_or_die(T->slots[i], new_size);

        ++T->m;
        ins_key(T->slots[i] + T->slot_sizes[i], key, len, &val);
        T->slot_sizes[i] = new_size;

        return val;
    }
    else return NULL;
}


value_t* ahtable_get(ahtable_t* T, const char* key, size_t len)
{
    return get_key(T, key, len, true);
}


value_t* ahtable_tryget(ahtable_t* T, const char* key, size_t len )
{
    return get_key(T, key, len, false);
}


int ahtable_del(ahtable_t* T, const char* key, size_t len)
{
    uint32_t i = hash(key, len) % T->n;
    size_t k;
    slot_t s;

    /* search the array for our key */
    s = T->slots[i];
    while ((size_t) (s - T->slots[i]) < T->slot_sizes[i]) {
        /* get the key length */
        k = keylen(s);
        s += k < 128 ? 1 : 2;

        /* skip keys that are longer than ours */
        if (k != len) {
            s += k + sizeof(value_t);
            continue;
        }

        /* key found. */
        if (memcmp(s, key, len) == 0) {
            /* move everything over, resize the array */
            unsigned char* t = s + len + sizeof(value_t);
            s -= k < 128 ? 1 : 2;
            memmove(s, t, T->slot_sizes[i] - (size_t) (t - T->slots[i]));
            T->slot_sizes[i] -= (size_t) (t - s);
            --T->m;
            return 0;
        }
        /* key not found. */
        else {
            s += k + sizeof(value_t);
            continue;
        }
    }

    // Key was not found. Do nothing.
    return -1;
}



static int cmpkey(const void* a_, const void* b_)
{
    slot_t a = *(slot_t*) a_;
    slot_t b = *(slot_t*) b_;

    size_t ka = keylen(a), kb = keylen(b);

    a += ka < 128 ? 1 : 2;
    b += kb < 128 ? 1 : 2;

    int c = memcmp(a, b, ka < kb ? ka : kb);
    return c == 0 ? (int) ka - (int) kb : c;
}


/* Sorted/unsorted iterators are kept private and exposed by passing the
sorted flag to ahtable_iter_begin. */

typedef struct ahtable_sorted_iter_t_
{
    const ahtable_t* T; // parent
    slot_t* xs; // pointers to keys
    size_t i; // current key
} ahtable_sorted_iter_t;


static ahtable_sorted_iter_t* ahtable_sorted_iter_begin(const ahtable_t* T)
{
    ahtable_sorted_iter_t* i = (ahtable_sorted_iter_t*) malloc_or_die(sizeof(ahtable_sorted_iter_t));
    i->T = T;
    i->xs = (slot_t *) malloc_or_die(T->m * sizeof(slot_t));
    i->i = 0;

    slot_t s;
    size_t j, k, u;
    for (j = 0, u = 0; j < T->n; ++j) {
        s = T->slots[j];
        while (s < T->slots[j] + T->slot_sizes[j]) {
            i->xs[u++] = s;
            k = keylen(s);
            s += k < 128 ? 1 : 2;
            s += k + sizeof(value_t);
        }
    }

    qsort(i->xs, T->m, sizeof(slot_t), cmpkey);

    return i;
}


static bool ahtable_sorted_iter_finished(ahtable_sorted_iter_t* i)
{
    return i->i >= i->T->m;
}


static void ahtable_sorted_iter_next(ahtable_sorted_iter_t* i)
{
    if (ahtable_sorted_iter_finished(i)) return;
    ++i->i;
}


static void ahtable_sorted_iter_free(ahtable_sorted_iter_t* i)
{
    if (i == NULL) return;
    free(i->xs);
    free(i);
}


static const char* ahtable_sorted_iter_key(ahtable_sorted_iter_t* i, size_t* len)
{
    if (ahtable_sorted_iter_finished(i)) return NULL;

    slot_t s = i->xs[i->i];
    *len = keylen(s);

    return (const char*) (s + (*len < 128 ? 1 : 2));
}


static value_t*  ahtable_sorted_iter_val(ahtable_sorted_iter_t* i)
{
    if (ahtable_sorted_iter_finished(i)) return NULL;

    slot_t s = i->xs[i->i];
    size_t k = keylen(s);

    s += k < 128 ? 1 : 2;
    s += k;

    return (value_t*) s;
}


typedef struct ahtable_unsorted_iter_t_
{
    const ahtable_t* T; // parent
    size_t i;           // slot index
    slot_t s;           // slot position
} ahtable_unsorted_iter_t;


static ahtable_unsorted_iter_t* ahtable_unsorted_iter_begin(const ahtable_t* T)
{
    ahtable_unsorted_iter_t* i = (ahtable_unsorted_iter_t*) malloc_or_die(sizeof(ahtable_unsorted_iter_t));
    i->T = T;

    for (i->i = 0; i->i < i->T->n; ++i->i) {
        i->s = T->slots[i->i];
        if ((size_t) (i->s - T->slots[i->i]) >= T->slot_sizes[i->i]) continue;
        break;
    }

    return i;
}


static bool ahtable_unsorted_iter_finished(ahtable_unsorted_iter_t* i)
{
    return i->i >= i->T->n;
}


static void ahtable_unsorted_iter_next(ahtable_unsorted_iter_t* i)
{
    if (ahtable_unsorted_iter_finished(i)) return;

    /* get the key length */
    size_t k = keylen(i->s);
    i->s += k < 128 ? 1 : 2;

    /* skip to the next key */
    i->s += k + sizeof(value_t);

    if ((size_t) (i->s - i->T->slots[i->i]) >= i->T->slot_sizes[i->i]) {
        do {
            ++i->i;
        } while(i->i < i->T->n &&
                i->T->slot_sizes[i->i] == 0);

        if (i->i < i->T->n) i->s = i->T->slots[i->i];
        else i->s = NULL;
    }
}


static void ahtable_unsorted_iter_free(ahtable_unsorted_iter_t* i)
{
    free(i);
}


static const char* ahtable_unsorted_iter_key(ahtable_unsorted_iter_t* i, size_t* len)
{
    if (ahtable_unsorted_iter_finished(i)) return NULL;

    slot_t s = i->s;
    size_t k;
    if (0x1 & *s) {
        k = (size_t) (*((uint16_t*) s)) >> 1;
        s += 2;
    }
    else {
        k = (size_t) (*s >> 1);
        s += 1;
    }

    *len = k;
    return (const char*) s;
}


static value_t* ahtable_unsorted_iter_val(ahtable_unsorted_iter_t* i)
{
    if (ahtable_unsorted_iter_finished(i)) return NULL;

    slot_t s = i->s;

    size_t k;
    if (0x1 & *s) {
        k = (size_t) (*((uint16_t*) s)) >> 1;
        s += 2;
    }
    else {
        k = (size_t) (*s >> 1);
        s += 1;
    }

    s += k;
    return (value_t*) s;
}


struct ahtable_iter_t_
{
    bool sorted;
    union {
        ahtable_unsorted_iter_t* unsorted;
        ahtable_sorted_iter_t* sorted;
    } i;
};


ahtable_iter_t* ahtable_iter_begin(const ahtable_t* T, bool sorted) {
    ahtable_iter_t* i = (ahtable_iter_t *) malloc_or_die(sizeof(ahtable_iter_t));
    i->sorted = sorted;
    if (sorted) i->i.sorted   = ahtable_sorted_iter_begin(T);
    else        i->i.unsorted = ahtable_unsorted_iter_begin(T);
    return i;
}


void ahtable_iter_next(ahtable_iter_t* i)
{
    if (i->sorted) ahtable_sorted_iter_next(i->i.sorted);
    else           ahtable_unsorted_iter_next(i->i.unsorted);
}


bool ahtable_iter_finished(ahtable_iter_t* i)
{
    if (i->sorted) return ahtable_sorted_iter_finished(i->i.sorted);
    else           return ahtable_unsorted_iter_finished(i->i.unsorted);
}


void ahtable_iter_free(ahtable_iter_t* i)
{
    if (i == NULL) return;
    if (i->sorted) ahtable_sorted_iter_free(i->i.sorted);
    else           ahtable_unsorted_iter_free(i->i.unsorted);
    free(i);
}


const char* ahtable_iter_key(ahtable_iter_t* i, size_t* len)
{
    if (i->sorted) return ahtable_sorted_iter_key(i->i.sorted, len);
    else           return ahtable_unsorted_iter_key(i->i.unsorted, len);
}


value_t* ahtable_iter_val(ahtable_iter_t* i)
{
    if (i->sorted) return ahtable_sorted_iter_val(i->i.sorted);
    else           return ahtable_unsorted_iter_val(i->i.unsorted);
}


void* malloc_or_die(size_t n)
{
    void* p = malloc(n);
    if (p == NULL && n != 0) {
        fprintf(stderr, "Cannot allocate %zu bytes.\n", n);
        exit(EXIT_FAILURE);
    }
    return p;
}


void* realloc_or_die(void* ptr, size_t n)
{
    void* p = realloc(ptr, n);
    if (p == NULL && n != 0) {
        fprintf(stderr, "Cannot allocate %zu bytes.\n", n);
        exit(EXIT_FAILURE);
    }
    return p;
}


FILE* fopen_or_die(const char* path, const char* mode)
{
    FILE* f = fopen(path, mode);
    if (f == NULL) {
        fprintf(stderr, "Cannot open file %s with mode %s.\n", path, mode);
        exit(EXIT_FAILURE);
    }
    return f;
}


#include <assert.h>
#include <string.h>

#define HT_UNUSED(x) x=x

/* maximum number of keys that may be stored in a bucket before it is burst */
static const size_t MAX_BUCKET_SIZE = 16384;
#define NODE_MAXCHAR 0xff // 0x7f for 7-bit ASCII
#define NODE_CHILDS (NODE_MAXCHAR+1)

static const uint8_t NODE_TYPE_TRIE          = 0x1;
static const uint8_t NODE_TYPE_PURE_BUCKET   = 0x2;
static const uint8_t NODE_TYPE_HYBRID_BUCKET = 0x4;
static const uint8_t NODE_HAS_VAL            = 0x8;


struct trie_node_t_;

/* Node's may be trie nodes or buckets. This union allows us to keep
 * non-specific pointer. */
typedef union node_ptr_
{
    ahtable_t*           b;
    struct trie_node_t_* t;
    uint8_t*             flag;
} node_ptr;


typedef struct trie_node_t_
{
    uint8_t flag;

    /* the value for the key that is consumed on a trie node */
    value_t val;

    /* Map a character to either a trie_node_t or a ahtable_t. The first byte
     * must be examined to determine which. */
    node_ptr xs[NODE_CHILDS];

} trie_node_t;

struct hattrie_t_
{
    node_ptr root; // root node
    size_t m;      // number of stored keys
};

/* Create a new trie node with all pointer pointing to the given child (which
 * can be NULL). */
static trie_node_t* alloc_trie_node(hattrie_t* T, node_ptr child)
{
    trie_node_t* node = (trie_node_t *) malloc_or_die(sizeof(trie_node_t));
    node->flag = NODE_TYPE_TRIE;
    node->val  = 0;

    /* pass T to allow custom allocator for trie. */
    HT_UNUSED(T); /* unused now */

    size_t i;
    for (i = 0; i < NODE_CHILDS; ++i) node->xs[i] = child;
    return node;
}

/* iterate trie nodes until string is consumed or bucket is found */
static node_ptr hattrie_consume(node_ptr *p, const char **k, size_t *l, unsigned brk)
{
    node_ptr node = p->t->xs[(unsigned char) **k];
    while (*node.flag & NODE_TYPE_TRIE && *l > brk) {
        ++*k;
        --*l;
        *p   = node;
        node = node.t->xs[(unsigned char) **k];
    }

    /* copy and writeback variables if it's faster */

    assert(*p->flag & NODE_TYPE_TRIE);
    return node;
}

/* use node value and return pointer to it */
static inline value_t* hattrie_useval(hattrie_t *T, node_ptr n)
{
    if (!(n.t->flag & NODE_HAS_VAL)) {
        n.t->flag |= NODE_HAS_VAL;
        ++T->m;
    }
    return &n.t->val;
}

/* clear node value if exists */
static inline int hattrie_clrval(hattrie_t *T, node_ptr n)
{
    if (n.t->flag & NODE_HAS_VAL) {
        n.t->flag &= ~NODE_HAS_VAL;
        n.t->val = 0;
        --T->m;
        return 0;
    }
    return -1;
}

/* find node in trie */
static node_ptr hattrie_find(hattrie_t* T, const char **key, size_t *len, int* found)
{
    *found = 1;
    node_ptr parent = T->root;
    assert(*parent.flag & NODE_TYPE_TRIE);

    if (*len == 0) return parent;

    node_ptr node = hattrie_consume(&parent, key, len, 1);

    /* if the trie node consumes value, use it */
    if (*node.flag & NODE_TYPE_TRIE) {
        if (!(node.t->flag & NODE_HAS_VAL)) {
            *found = 0;
        }
        *key += 1;
        *len -= 1;
        return node;
    }

    /* pure bucket holds only key suffixes, skip current char */
    if (*node.flag & NODE_TYPE_PURE_BUCKET) {
        *key += 1;
        *len -= 1;
    }

    /* do not scan bucket, it's not needed for this operation */
    return node;
}

hattrie_t* hattrie_create()
{
    hattrie_t* T = (hattrie_t *) malloc_or_die(sizeof(hattrie_t));
    T->m = 0;

    node_ptr node;
    node.b = ahtable_create();
    node.b->flag = NODE_TYPE_HYBRID_BUCKET;
    node.b->c0 = 0x00;
    node.b->c1 = NODE_MAXCHAR;
    T->root.t = alloc_trie_node(T, node);

    return T;
}


static void hattrie_free_node(node_ptr node)
{
    if (*node.flag & NODE_TYPE_TRIE) {
        size_t i;
        for (i = 0; i < NODE_CHILDS; ++i) {
            if (i > 0 && node.t->xs[i].t == node.t->xs[i - 1].t) continue;

            /* XXX: recursion might not be the best choice here. It is possible
             * to build a very deep trie. */
            if (node.t->xs[i].t) hattrie_free_node(node.t->xs[i]);
        }
        free(node.t);
    }
    else {
        ahtable_free(node.b);
    }
}


void hattrie_free(hattrie_t* T)
{
    hattrie_free_node(T->root);
    free(T);
}


size_t hattrie_size(hattrie_t* T)
{
    return T->m;
}


/* Perform one split operation on the given node with the given parent.
 */
static void hattrie_split(hattrie_t* T, node_ptr parent, node_ptr node)
{
    /* only buckets may be split */
    assert(*node.flag & NODE_TYPE_PURE_BUCKET ||
           *node.flag & NODE_TYPE_HYBRID_BUCKET);

    assert(*parent.flag & NODE_TYPE_TRIE);

    if (*node.flag & NODE_TYPE_PURE_BUCKET) {
        /* turn the pure bucket into a hybrid bucket */
        parent.t->xs[node.b->c0].t = alloc_trie_node(T, node);

        /* if the bucket had an empty key, move it to the new trie node */
        value_t* val = ahtable_tryget(node.b, NULL, 0);
        if (val) {
            parent.t->xs[node.b->c0].t->val     = *val;
            parent.t->xs[node.b->c0].t->flag |= NODE_HAS_VAL;
            *val = 0;
            ahtable_del(node.b, NULL, 0);
        }

        node.b->c0   = 0x00;
        node.b->c1   = NODE_MAXCHAR;
        node.b->flag = NODE_TYPE_HYBRID_BUCKET;

        return;
    }

    /* This is a hybrid bucket. Perform a proper split. */

    /* count the number of occourances of every leading character */
    unsigned int cs[NODE_CHILDS]; // occurance count for leading chars
    memset(cs, 0, NODE_CHILDS * sizeof(unsigned int));
    size_t len;
    const char* key;

    ahtable_iter_t* i = ahtable_iter_begin(node.b, false);
    while (!ahtable_iter_finished(i)) {
        key = ahtable_iter_key(i, &len);
        assert(len > 0);
        cs[(unsigned char) key[0]] += 1;
        ahtable_iter_next(i);
    }
    ahtable_iter_free(i);

    /* choose a split point */
    unsigned int left_m, right_m, all_m;
    unsigned char j = node.b->c0;
    all_m   = ahtable_size(node.b);
    left_m  = cs[j];
    right_m = all_m - left_m;
    int d;

    while (j + 1 < node.b->c1) {
        d = abs((int) (left_m + cs[j + 1]) - (int) (right_m - cs[j + 1]));
        if (d <= abs((int) (left_m - right_m)) && left_m + cs[j + 1] < all_m) {
            j += 1;
            left_m  += cs[j];
            right_m -= cs[j];
        }
        else break;
    }

    /* now split into two node cooresponding to ranges [0, j] and
     * [j + 1, NODE_MAXCHAR], respectively. */


    /* create new left and right nodes */

    /* TODO: Add a special case if either node is a hybrid bucket containing all
     * the keys. In such a case, do not build a new table, just use the old one.
     * */
    size_t num_slots;


    for (num_slots = ahtable_initial_size;
            (double) left_m > ahtable_max_load_factor * (double) num_slots;
            num_slots *= 2);

    node_ptr left, right;
    left.b  = ahtable_create_n(num_slots);
    left.b->c0   = node.b->c0;
    left.b->c1   = j;
    left.b->flag = left.b->c0 == left.b->c1 ?
                      NODE_TYPE_PURE_BUCKET : NODE_TYPE_HYBRID_BUCKET;


    for (num_slots = ahtable_initial_size;
            (double) right_m > ahtable_max_load_factor * (double) num_slots;
            num_slots *= 2);

    right.b = ahtable_create_n(num_slots);
    right.b->c0   = j + 1;
    right.b->c1   = node.b->c1;
    right.b->flag = right.b->c0 == right.b->c1 ?
                      NODE_TYPE_PURE_BUCKET : NODE_TYPE_HYBRID_BUCKET;


    /* update the parent's pointer */

    unsigned int c;
    for (c = node.b->c0; c <= j; ++c) parent.t->xs[c] = left;
    for (; c <= node.b->c1; ++c)      parent.t->xs[c] = right;



    /* distribute keys to the new left or right node */
    value_t* u;
    value_t* v;
    i = ahtable_iter_begin(node.b, false);
    while (!ahtable_iter_finished(i)) {
        key = ahtable_iter_key(i, &len);
        u   = ahtable_iter_val(i);
        assert(len > 0);

        /* left */
        if ((unsigned char) key[0] <= j) {
            if (*left.flag & NODE_TYPE_PURE_BUCKET) {
                v = ahtable_get(left.b, key + 1, len - 1);
            }
            else {
                v = ahtable_get(left.b, key, len);
            }
            *v = *u;
        }

        /* right */
        else {
            if (*right.flag & NODE_TYPE_PURE_BUCKET) {
                v = ahtable_get(right.b, key + 1, len - 1);
            }
            else {
                v = ahtable_get(right.b, key, len);
            }
            *v = *u;
        }

        ahtable_iter_next(i);
    }

    ahtable_iter_free(i);
    ahtable_free(node.b);
}

value_t* hattrie_get(hattrie_t* T, const char* key, size_t len)
{
    node_ptr parent = T->root;
    assert(*parent.flag & NODE_TYPE_TRIE);

    if (len == 0) return &parent.t->val;

    /* consume all trie nodes, now parent must be trie and child anything */
    node_ptr node = hattrie_consume(&parent, &key, &len, 0);
    assert(*parent.flag & NODE_TYPE_TRIE);

    /* if the key has been consumed on a trie node, use its value */
    if (len == 0) {
        if (*node.flag & NODE_TYPE_TRIE) {
            return hattrie_useval(T, node);
        }
        else if (*node.flag & NODE_TYPE_HYBRID_BUCKET) {
            return hattrie_useval(T, parent);
        }
    }


    /* preemptively split the bucket if it is full */
    while (ahtable_size(node.b) >= MAX_BUCKET_SIZE) {
        hattrie_split(T, parent, node);

        /* after the split, the node pointer is invalidated, so we search from
         * the parent again. */
        node = hattrie_consume(&parent, &key, &len, 0);

        /* if the key has been consumed on a trie node, use its value */
        if (len == 0) {
            if (*node.flag & NODE_TYPE_TRIE) {
                return hattrie_useval(T, node);
            }
            else if (*node.flag & NODE_TYPE_HYBRID_BUCKET) {
                return hattrie_useval(T, parent);
            }
        }
    }

    assert(*node.flag & NODE_TYPE_PURE_BUCKET || *node.flag & NODE_TYPE_HYBRID_BUCKET);

    assert(len > 0);
    size_t m_old = node.b->m;
    value_t* val;
    if (*node.flag & NODE_TYPE_PURE_BUCKET) {
        val = ahtable_get(node.b, key + 1, len - 1);
    }
    else {
        val = ahtable_get(node.b, key, len);
    }
    T->m += (node.b->m - m_old);

    return val;
}


value_t* hattrie_tryget(hattrie_t* T, const char* key, size_t len)
{
    /* find node for given key */
    int found;
    node_ptr node = hattrie_find(T, &key, &len, &found);
    if (!found) {
        return NULL;
    }

    /* if the trie node consumes value, use it */
    if (*node.flag & NODE_TYPE_TRIE) {
        return &node.t->val;
    }

    return ahtable_tryget(node.b, key, len);
}


void hattrie_walk (hattrie_t* T, const char* key, size_t len, void* user_data, hattrie_walk_cb cb) {
    unsigned char* k = (unsigned char*)key;
    node_ptr node = T->root;
    size_t i, j;
    ahtable_iter_t* it;

    /* go down until a bucket is reached */
    for (i = 0; i < len; i++, k++) {
        if (!(*node.flag & NODE_TYPE_TRIE))
            break;
        node = node.t->xs[*k];
        if (*node.flag & NODE_HAS_VAL) {
            if (hattrie_walk_stop == cb(key, i, &node.t->val, user_data))
                return;
        }
    }
    if (i == len)
        return;

    assert(i);
    if (*node.flag & NODE_TYPE_HYBRID_BUCKET) {
        i--;
        k--;
    } else {
        assert(*node.flag & NODE_TYPE_PURE_BUCKET);
    }

    /* dict order ensured short => long */
    it = ahtable_iter_begin(node.b, true);
    for(; !ahtable_iter_finished(it); ahtable_iter_next(it)) {
        size_t stored_len;
        unsigned char* stored_key = (unsigned char*)ahtable_iter_key(it, &stored_len);
        int matched = 1;
        if (stored_len + i > len) {
            continue;
        }
        for (j = 0; j < stored_len; j++) {
            if (stored_key[j] != k[j]) {
                matched = 0;
                break;
            }
        }
        if (matched) {
            value_t* val = ahtable_iter_val(it);
            if (hattrie_walk_stop == cb(key, i + stored_len, val, user_data)) {
                ahtable_iter_free(it);
                return;
            }
        }
    }
    ahtable_iter_free(it);
}


int hattrie_del(hattrie_t* T, const char* key, size_t len)
{
    node_ptr parent = T->root;
    assert(*parent.flag & NODE_TYPE_TRIE);

    /* find node for deletion */
    int found;
    node_ptr node = hattrie_find(T, &key, &len, &found);
    if (!found) {
        return -1;
    }

    /* if consumed on a trie node, clear the value */
    if (*node.flag & NODE_TYPE_TRIE) {
        return hattrie_clrval(T, node);
    }

    /* remove from bucket */
    size_t m_old = ahtable_size(node.b);
    int ret =  ahtable_del(node.b, key, len);
    T->m -= (m_old - ahtable_size(node.b));

    /* merge empty buckets */
    /*! \todo */

    return ret;
}


/* plan for iteration:
 * This is tricky, as we have no parent pointers currently, and I would like to
 * avoid adding them. That means maintaining a stack
 *
 */

typedef struct hattrie_node_stack_t_
{
    unsigned char   c;
    size_t level;

    node_ptr node;
    struct hattrie_node_stack_t_* next;

} hattrie_node_stack_t;


struct hattrie_iter_t_
{
    char* key;
    size_t keysize; // space reserved for the key
    size_t level;

    /* keep track of keys stored in trie nodes */
    bool    has_nil_key;
    value_t nil_val;

    const hattrie_t* T;
    bool sorted;
    ahtable_iter_t* i;
    hattrie_node_stack_t* stack;

    // subtree inside a table
    // store remaining prefix for filtering nodes not matching it
    char* prefix;
    size_t prefix_len;
};


static void hattrie_iter_pushchar(hattrie_iter_t* i, size_t level, char c)
{
    if (i->keysize < level) {
        i->keysize *= 2;
        i->key = (char *) realloc_or_die(i->key, i->keysize * sizeof(char));
    }

    if (level > 0) {
        i->key[level - 1] = c;
    }

    i->level = level;
}


static void hattrie_iter_nextnode(hattrie_iter_t* i)
{
    if (i->stack == NULL) return;

    /* pop the stack */
    node_ptr node;
    hattrie_node_stack_t* next;
    unsigned char   c;
    size_t level;

    node  = i->stack->node;
    next  = i->stack->next;
    c     = i->stack->c;
    level = i->stack->level;

    free(i->stack);
    i->stack = next;

    if (*node.flag & NODE_TYPE_TRIE) {
        hattrie_iter_pushchar(i, level, c);

        if(node.t->flag & NODE_HAS_VAL) {
            i->has_nil_key = true;
            i->nil_val = node.t->val;
        }

        /* push all child nodes from right to left */
        int j;
        for (j = NODE_MAXCHAR; j >= 0; --j) {

            /* skip repeated pointers to hybrid bucket */
            if (j < NODE_MAXCHAR && node.t->xs[j].t == node.t->xs[j + 1].t) continue;

            // push stack
            next = i->stack;
            i->stack = (hattrie_node_stack_t *) malloc_or_die(sizeof(hattrie_node_stack_t));
            i->stack->node  = node.t->xs[j];
            i->stack->next  = next;
            i->stack->level = level + 1;
            i->stack->c     = (unsigned char) j;
        }
    }
    else {
        if (*node.flag & NODE_TYPE_PURE_BUCKET) {
            hattrie_iter_pushchar(i, level, c);
        }
        else if (level) {
            i->level = level - 1;
        }
        i->i = ahtable_iter_begin(node.b, i->sorted);

    }
}


/** next non-nil-key node
 * TODO pick a better name
 */
static void hattrie_iter_step(hattrie_iter_t* i)
{
    while (((i->i == NULL || ahtable_iter_finished(i->i)) && !i->has_nil_key) &&
           i->stack != NULL ) {

        ahtable_iter_free(i->i);
        i->i = NULL;
        hattrie_iter_nextnode(i);
    }

    if (i->i != NULL && ahtable_iter_finished(i->i)) {
        ahtable_iter_free(i->i);
        i->i = NULL;
    }
}

static bool hattrie_iter_prefix_not_match(hattrie_iter_t* i)
{
    if (hattrie_iter_finished(i)) {
        return false; // can not advance the iter
    }
    if (i->level >= i->prefix_len) {
        return memcmp(i->key, i->prefix, i->prefix_len);
    } else if (i->has_nil_key) {
        return true; // subkey too short
    }

    size_t sublen;
    const char* subkey;
    subkey = ahtable_iter_key(i->i, &sublen);
    if (i->level + sublen < i->prefix_len) {
        return true; // subkey too short
    }
    return memcmp(i->key, i->prefix, i->level) ||
           memcmp(subkey, i->prefix + i->level, (i->prefix_len - i->level));
}


hattrie_iter_t* hattrie_iter_begin(const hattrie_t* T, bool sorted)
{
	return (hattrie_iter_t*) hattrie_iter_with_prefix(T, sorted, NULL, 0);
}


hattrie_iter_t* hattrie_iter_with_prefix(const hattrie_t* T, bool sorted, const char* prefix, size_t prefix_len)
{
    int found;
    node_ptr node = hattrie_find((hattrie_t*)T, &prefix, &prefix_len, &found);

    hattrie_iter_t* i = (hattrie_iter_t *) malloc_or_die(sizeof(hattrie_iter_t));
    i->T = T;
    i->sorted = sorted;
    i->i = NULL;
    i->keysize = 16;
    i->key = (char *) malloc_or_die(i->keysize * sizeof(char));
    i->level   = 0;
    i->has_nil_key = false;
    i->nil_val     = 0;

    i->prefix_len  = prefix_len;
    if (prefix_len) {
        i->prefix  = (char*)malloc_or_die(prefix_len);
        memcpy(i->prefix, prefix, prefix_len);
    } else {
        i->prefix  = NULL;
    }

    i->stack = (hattrie_node_stack_t *) malloc_or_die(sizeof(hattrie_node_stack_t));
    i->stack->next   = NULL;
    i->stack->node   = node;
    i->stack->c      = '\0';
    i->stack->level  = 0;

    hattrie_iter_step(i);
    if (i->prefix_len && hattrie_iter_prefix_not_match(i)) {
        hattrie_iter_next(i);
    }

    return i;
}


void hattrie_iter_next(hattrie_iter_t* i)
{
    do {
        if (hattrie_iter_finished(i)) return;

        if (i->i != NULL && !ahtable_iter_finished(i->i)) {
            ahtable_iter_next(i->i);
        }
        else if (i->has_nil_key) {
            i->has_nil_key = false;
            i->nil_val = 0;
            hattrie_iter_nextnode(i);
        }

        hattrie_iter_step(i);
    } while (i->prefix_len && hattrie_iter_prefix_not_match(i));
}


bool hattrie_iter_finished(hattrie_iter_t* i)
{
    return i->stack == NULL && i->i == NULL && !i->has_nil_key;
}


void hattrie_iter_free(hattrie_iter_t* i)
{
    if (i == NULL) return;
    if (i->i) ahtable_iter_free(i->i);

    hattrie_node_stack_t* next;
    while (i->stack) {
        next = i->stack->next;
        free(i->stack);
        i->stack = next;
    }

    if (i->prefix_len) {
        free(i->prefix);
    }

    free(i->key);
    free(i);
}


const char* hattrie_iter_key(hattrie_iter_t* i, size_t* len)
{
    if (hattrie_iter_finished(i)) return NULL;

    size_t sublen;
    const char* subkey;

    if (i->has_nil_key) {
        subkey = NULL;
        sublen = 0;
    }
    else subkey = ahtable_iter_key(i->i, &sublen);

    if (i->keysize < i->level + sublen + 1) {
        while (i->keysize < i->level + sublen + 1) i->keysize *= 2;
        i->key = (char *) realloc_or_die(i->key, i->keysize * sizeof(char));
    }

    memcpy(i->key + i->level, subkey, sublen);
    i->key[i->level + sublen] = '\0';

    *len = i->level + sublen - i->prefix_len;
    return i->key + i->prefix_len;
}


value_t* hattrie_iter_val(hattrie_iter_t* i)
{
    if (i->has_nil_key) return &i->nil_val;

    if (hattrie_iter_finished(i)) return NULL;

    return ahtable_iter_val(i->i);
}