package io.icednut.coding_training

import java.time.LocalDate

import scala.io.StdIn
import scala.util.Try

object AppMain extends App {

  private val age = Try(StdIn.readLine("What is your current age? ")).map(_.toInt).getOrElse(0)
  private val retireAge = Try(StdIn.readLine("At what age would you like to retire? ")).map(_.toInt).getOrElse(0)
  private val thisYear = LocalDate.now().getYear
  private val remain: Int = retireAge - age

  println(s"You have $remain years left until you can retire.")
  println(s"It's $thisYear, so you can retire in ${thisYear + remain}.")
}
