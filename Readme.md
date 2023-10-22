## CSC482 - Software Deployment

Skills: GoLang, Docker, AWS(EC2, DynamoDB, IAM, Cloudwatch/trail)

<b><ins>Accomplishments!</ins></b>

<ol>
<li>Built simple <ins>GoLang</ins> polling agent to query stock data from Polygon.io.</li>
<li>Associated polling agent with a timer allowing for a continuous information pipeline.</li>
<li>Containerized('dockerized') program into image with <ins>Docker</ins>.</li>
<li>Converted original Docker image into a <ins>multi-stage Docker image</ins> for efficient resource utilization. </li>
<li>Decided against Polygon's API due to it's query limitations -> <ins>web scraped</ins> instead</li>
<li>Wrote scraped data into <ins>AWS DynamoDB</ins></li>
<li>Utilized Solarwind's <ins>Loggly</ins> utility to monitor my agent once it's deployed on EC2</li>
</ol>
