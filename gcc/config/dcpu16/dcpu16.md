;; Machine description for Moxie
;; Copyright (C) 2009 Free Software Foundation, Inc.
;; Contributed by Anthony Green <green@moxielogic.com>

;; This file is part of GCC.

;; GCC is free software; you can redistribute it and/or modify it
;; under the terms of the GNU General Public License as published
;; by the Free Software Foundation; either version 3, or (at your
;; option) any later version.

;; GCC is distributed in the hope that it will be useful, but WITHOUT
;; ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
;; or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public
;; License for more details.

;; You should have received a copy of the GNU General Public License
;; along with GCC; see the file COPYING3.  If not see
;; <http://www.gnu.org/licenses/>.

;; -------------------------------------------------------------------------
;; DCPU16 specific constraints, predicates and attributes
;; -------------------------------------------------------------------------

(define_insn "nop"
  [(const_int 0)]
  ""
  "set 0, 0")

(define_constraint "A"
  "An absolute address."
  (and (match_code "mem")
       (ior (match_test "GET_CODE (XEXP (op, 0)) == SYMBOL_REF")
	    (match_test "GET_CODE (XEXP (op, 0)) == LABEL_REF")
	    (match_test "GET_CODE (XEXP (op, 0)) == CONST"))))



(define_insn "jump"
[(set (pc)
(label_ref (match_operand 0 "" "A"))
)]
""
"set PC,  %l0"
)

(define_insn "indirect_jump"
 [(set (pc) (match_operand:SI 0 "nonimmediate_operand" "r"))]
""
"set PC,  %0"
)

(define_expand "movsi"
[(set (match_operand:SI 0 "nonimmediate_operand" "")
(match_operand:SI 1 "general_operand" "")
)]
""
{
if(GET_CODE(operands[0])==MEM && GET_CODE(operands[1])!=REG)
{
if(can_create_pseudo_p())
{
operands[1]=force_reg(SImode,operands[1]);
}}}
)

(define_insn "epilogue"
[(clobber (const_int 0))]
""
"set PC, POP"
)

(define_insn "IITB_return"
[(return)]
""
"jr \\$ra"
)