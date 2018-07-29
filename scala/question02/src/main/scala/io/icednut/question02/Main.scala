package io.icednut.question02

import scala.io.StdIn

object Main extends App {

  val input: String = StdIn.readLine("What is the input string? ")

  input match {
    case i if i == null || i.trim.isEmpty => println("input something!")
    case _ => print(s"$input has ${input.length} characters.")
  }
}
