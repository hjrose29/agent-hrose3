# CSC482 - Software Deployment - Agent Side
## Go [here](https://github.com/hjrose29/server-hrose3) for the server implementation!

Skills: GoLang, Docker, AWS(EC2, DynamoDB, IAM, Cloudwatch/trail), Loggly, Web Scraping, RESTful API, JSON

<b><ins>What I did!</ins></b>

<ol>
<li>Built simple <ins>GoLang</ins> polling agent to query stock data from Polygon.io.</li>
<li>Associated polling agent with a timer allowing for a continuous information pipeline.</li>
<li>Containerized('dockerized') program into image with <ins>Docker</ins>.</li>
<li>Converted original Docker image into a <ins>multi-stage Docker image</ins> for efficient resource utilization. </li>
<li>Decided against Polygon's API due to it's query limitations -> <ins>web scraped</ins> instead.</li>
<li>Wrote scraped data into <ins>AWS DynamoDB.</ins></li>
<li>Utilized Solarwind's <ins>Loggly</ins> utility to monitor my agent once it's deployed on EC2.</li>
<li>Began running agent on <ins>AWS Fargate</ins> instance.</li>
</ol>

<br>

<b><ins>Building Container</ins>:<b>
docker build . -t <IMAGE_NAME>

<b><ins>Running Container</ins>:<b><br>
docker run -e LOGGLY_TOKEN=<LOGGLY_TOKEN> -e AWS_ACCESS_KEY_ID=<SECRET_TOKEN1> -e AWS_SECRET_ACCESS_KEY=<SECRET_TOKEN2> -e AWS_DEFAULT_REGION=<REGION> <IMAGE_NAME>
xq