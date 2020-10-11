package io.icednut.coding.training.question05

import java.time.LocalDateTime

import scala.io.StdIn
import scala.util.Try

object MainApp extends App {

  private val ageStr = StdIn.readLine("What is your current age?")
  private val retireStr = StdIn.readLine("At what age would you like to retire?")
  val age = Try(ageStr.toInt).toOption.getOrElse(0)
  val retireAge = Try(retireStr.toInt).toOption.getOrElse(0)
  val difference = retireAge - age
  val currentYear: Int = LocalDateTime.now().getYear

  println(s"You have $difference years left until you can retire.")
  println(s"It's $currentYear, so you can retire in ${currentYear + difference}")


}
