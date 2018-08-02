package io.icednut.question03

import scala.io.StdIn

object Main extends App {

  val quote: String = StdIn.readLine("What is the quote? ")
  val name: String = StdIn.readLine("Who said it? ")
  println(s"""${name} says, "${quote}"""")
}