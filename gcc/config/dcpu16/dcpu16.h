/* Target Definitions for moxie.
   Copyright (C) 2008, 2009, 2010  Free Software Foundation, Inc.
   Contributed by Anthony Green.

   This file is part of GCC.

   GCC is free software; you can redistribute it and/or modify it
   under the terms of the GNU General Public License as published
   by the Free Software Foundation; either version 3, or (at your
   option) any later version.

   GCC is distributed in the hope that it will be useful, but WITHOUT
   ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
   or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public
   License for more details.

   You should have received a copy of the GNU General Public License
   along with GCC; see the file COPYING3.  If not see
   <http://www.gnu.org/licenses/>.  */

#ifndef GCC_DCPU16_H
#define GCC_DCPU16_H

/* Run-time Target Specification */

#define TARGET_CPU_CPP_BUILTINS() \
  { \
    builtin_define_std ("dcpu16");		\
    builtin_define_std ("DCPU16");		\
  }

#define BYTES_BIG_ENDIAN 1

#define WORDS_BIG_ENDIAN 1

#define BITS_PER_UNIT 16

#define UNITS_PER_WORD 1

#define INT_TYPE_SIZE 16

#define SHORT_TYPE_SIZE 16

#define LONG_TYPE_SIZE 32

#define LONG_LONG_TYPE_SIZE 64

#define CHAR_TYPE_SIZE 16

#define FLOAT_TYPE_SIZE 32

#define DOUBLE_TYPE_SIZE 64

#define LONG_DOUBLE_TYPE_SIZE 64

#define PARM_BOUNDARY 16

#define STACK_BOUNDARY 16

#define FUNCTION_BOUNDARY 16

#define BIGGEST_ALIGNMENT 16

#define SIZE_TYPE "unsigned int"

#define PTRDIFF_TYPE "int"

#define INT16_TYPE "int"

#define INT32_TYPE "long int"

#define INT64_TYPE "long long int"

#define INT16_TYPE "unsigned int"

#define INT32_TYPE "unsigned long int"

#define INT64_TYPE "unsigned long long int"



#define FIRST_PSEUDO_REGISTER 10

#define REGISTER_NAMES { \
  "A", "B", "C", \
  "X", "Y", "Z", \
  "I", "J", "SP", "PC" }

#define FIXED_REGISTERS  { 0, 0, 0, \
                           0, 0, 0, \
                           0, 1, 1, 1 }

#define CALL_USED_REGISTERS { 0, 0, 0, \
                              0, 0, 0, \
                              0, 1, 1, 1 }

#define HARD_REGNO_NREGS(REGNO, MODE) \
((GET_MODE_SIZE (MODE) + UNITS_PER_WORD - 1) \
/ UNITS_PER_WORD)


enum reg_class
{
  NO_REGS,
  GENERAL_REGS,
  STACK_REGS,
  PC_REGS,
  ALL_REGS
};

#define REG_CLASS_NAMES { \
  "NO_REGS", \
  "GENERAL_REGS", \
  "STACK_REGS", \
  "PC_REGS", \
  "ALL_REGS" \
}


#define REG_CLASS_CONTENTS \
{ { 0x00000000 }, /* Empty */			   \
  { 0x000000FF }, /* A,B,C,X,Y,Z,I,J */ \
  { 0x00000100 }, /* SP */	                   \
  { 0x00000200 }, /* PC */                        \
  { 0x000003FF }  /* All registers */              \
}

#define STACK_POINTER_REGNUM 8

#define FRAME_POINTER_REGNUM 7

#define STACK_GROWS_DOWNWARD

#define STACK_ALIGNMENT_NEEDED 0

#define STACK_POINTER_OFFSET TODO

#define FIRST_PARAM_OFFSET TODO



#endif // GCC_DCPU16_H
