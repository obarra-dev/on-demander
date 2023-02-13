package app

import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.Job
import kotlinx.coroutines.joinAll
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody.Companion.toRequestBody

fun main(args: Array<String>) {
  val url = "http://localhost:3000/api/web-console/v1/ondemand/jobs"
  val cookie =
    "_ga=GA1.1.484072228.1662045000; _biz_uid=f06e970d51f44231a03f8e64a50e7f2b; _biz_flagsA=%7B%22Version%22%3A1%2C%22ViewThrough%22%3A%221%22%2C%22XDomain%22%3A%221%22%7D; hubspotutk=cbc533be50f7620d527227e2034176db; apt.uid=AP-S8VLIJURGKY0-2-4-1662045004869-98433362.0.2.965c0cd2-500b-424e-a457-f8b6c21d6e35; __hstc=181257784.cbc533be50f7620d527227e2034176db.1662045002166.1667244546928.1668181648383.17; _biz_nA=108; _biz_pendingA=%5B%22m%2Fipv%3F_biz_r%3D%26_biz_h%3D802059049%26_biz_u%3Df06e970d51f44231a03f8e64a50e7f2b%26_biz_s%3D5ac105%26_biz_l%3Dhttp%253A%252F%252Flocalhost%253A3000%252Fresults%252Fgithub.com%252Fobarra-dev%252Fminesweeper-API%252F01GK2H4MYYV5J06DQJT09X73X9%252F%253Fph%253DPhaseIntroduced%26_biz_t%3D1669753816188%26_biz_i%3DSonatype%2520Lift%2520--%2520Console%26_biz_n%3D107%26rnd%3D871882%22%5D; muserAuth=2aVWSduSQ1uwGbh8hAUGZUi96HjosSrvhwaxioSRwWSgS1FFte1ZCGUvuMw9BfphrR8hWQAb9c7cwS7BNh8FSbZcLfk9tJZm5HXGt3dzbWpmqrbV3hZF9JzZ5BRT9at3FNytjPcpKA8eZ4jxG5JQE55QA3T4yc63KAWR2vSNiqB1EUvRSP; muserSiteCookie=Ks5v1yoE6Xv1jRZKWk9TyYgT4YBZJAJfWMsbyyKqM98sC55j3TtF6gXuWayYQTnfppXUWgBwiMuMYFASpY4t4bTWJSthwyr5MgRUXTHF"

  val repos = getRepos()

  println("Count repos: ${repos.size}")
  val okHttpClient = OkHttpClient()

  runBlocking {
    val jobs: List<Job> = repos.map {
      launch(context = Dispatchers.Default) {
        val request = Request.Builder()
          .header("cookie", cookie)
          .header("Content-Type", "application/json")
          .method("POST", it.toRequestBody())
          .url(url)
          .build()

        val response = okHttpClient.newCall(request).execute()
        val result = response.body?.string()
        println("Res: $result, for: $it")
      }
    }
    println("[${Thread.currentThread().name}] Created all coroutines")
    jobs.joinAll()
    println("[${Thread.currentThread().name}] Finished all coroutines")
  }
}

fun getRepos(): List<String> {
  return listOf(
    """{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"add-errors","generatePrToFixIssues":false}""",
    """{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"main","generatePrToFixIssues":false}""",
    """{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"errorprone-2.18.0","generatePrToFixIssues":false}""",
    """{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"main","generatePrToFixIssues":false}""",
  )
}