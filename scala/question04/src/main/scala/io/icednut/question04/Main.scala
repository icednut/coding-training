package io.icednut.question04

import scala.io.StdIn

object Main extends App {

  val noun: String = StdIn.readLine("Enter an noun: ")
  val verb: String = StdIn.readLine("Enter a verb: ")
  val adjective: String = StdIn.readLine("Enter an adjective: ")
  val adverb: String = StdIn.readLine("Enter an adverb: ")

  println(
    s"""
       |Do you ${verb} your ${adjective} ${noun} ${adverb}? That's hilarious!
     """.stripMargin)

}
