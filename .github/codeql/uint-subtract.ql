/**
 * @name Avoid uint subtraction
 * @description Subtraction of uints can underflow and cause unexpected behavior.
 *              Prefer using `subtract.Uint64(a,b)` which return 0 if b is larger than a.
 * @kind problem
 * @problem.severity error
 * @id go/uint-subtract
 * @precision high
 * @tags correctness
 */

 import go

 from SubExpr ex
 where
    ex.getLeftOperand().getType().getName() = "uint64" and
    ex.getRightOperand().getType().getName() = "uint64"
 select ex, "Subtraction between two uint64s detected."
