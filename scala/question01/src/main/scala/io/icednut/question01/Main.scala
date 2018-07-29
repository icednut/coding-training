package io.icednut.question01

import scala.io.StdIn

object Main extends App {
  print("What is your name? ")
  println(s"Hello, ${StdIn.readLine()}, nice to meet you!")
}
