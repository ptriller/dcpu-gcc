/* InvalidAlgorithmParameterException.java -- an invalid parameter to a
   security algorithm
   Copyright (C) 2000, 2002, 2005, 2006  Free Software Foundation, Inc.

This file is part of GNU Classpath.

GNU Classpath is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2, or (at your option)
any later version.

GNU Classpath is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
General Public License for more details.

You should have received a copy of the GNU General Public License
along with GNU Classpath; see the file COPYING.  If not, write to the
Free Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
02110-1301 USA.

Linking this library statically or dynamically with other modules is
making a combined work based on this library.  Thus, the terms and
conditions of the GNU General Public License cover the whole
combination.

As a special exception, the copyright holders of this library give you
permission to link this library with independent modules to produce an
executable, regardless of the license terms of these independent
modules, and to copy and distribute the resulting executable under
terms of your choice, provided that you also meet, for each linked
independent module, the terms and conditions of the license of that
module.  An independent module is a module which is not derived from
or based on this library.  If you modify this library, you may extend
this exception to your version of the library, but you are not
obligated to do so.  If you do not wish to do so, delete this
exception statement from your version. */

package java.security;

/**
 * Thrown for an invalid security algorithm parameter.
 *
 * @author Warren Levy (warrenl@cygnus.com)
 * @since 1.2
 * @status updated to 1.4
 */
public class InvalidAlgorithmParameterException
  extends GeneralSecurityException
{
  /**
   * Compatible with JDK 1.2+.
   */
  private static final long serialVersionUID = 2864672297499471472L;

  /**
   * Construct an exception with no message.
   */
  public InvalidAlgorithmParameterException()
  {
    super();
  }

  /**
   * Construct an exception with a message.
   *
   * @param msg the message
   */
  public InvalidAlgorithmParameterException(String msg)
  {
    super(msg);
  }

  /**
   * Create a new instance with a descriptive error message and
   * a cause.
   * @param s the descriptive error message
   * @param cause the cause
   * @since 1.5
   */
  public InvalidAlgorithmParameterException(String s, Throwable cause)
  {
    super(s, cause);
  }

  /**
   * Create a new instance with a cause.
   * @param cause the cause
   * @since 1.5
   */
  public InvalidAlgorithmParameterException(Throwable cause)
  {
    super(cause);
  }
}
