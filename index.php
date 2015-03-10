<?php

function CallAPI($method, $url, $data = false)
{
    $curl = curl_init();

    switch ($method)
    {
        case "POST":
            curl_setopt($curl, CURLOPT_POST, 1);

            if ($data)
                curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
            break;
        case "PUT":
            curl_setopt($curl, CURLOPT_PUT, 1);
            break;
        default:
            if ($data)
                $url = sprintf("%s?%s", $url, http_build_query($data));
    }

    // Optional Authentication:
    curl_setopt($curl, CURLOPT_HTTPAUTH, CURLAUTH_BASIC);
    curl_setopt($curl, CURLOPT_USERPWD, "username:password");

    curl_setopt($curl, CURLOPT_URL, $url);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);

    $result = curl_exec($curl);

    curl_close($curl);

    return $result;
}

$response = CallAPI("GET", "localhost:8080/all");
$json = json_decode($response);

$albums = $json->ResultSet; 

// Output the search results
foreach ($json as $row) {

    echo "<span><h4>Title".$row->title."</h4> Author: ".$row->author." Category: ".$row->category."<br><b>Content: ".$row->content."</b></span><br>";
    
    unset($comment);
    foreach ($row->comments as $comment) {
    	echo "<div>". $comment."</div>";
    }

    echo "<form action=comment.php method=post>";
    echo "Comment: <input type=\"hidden\" id=\"postId\" name=\"postId\" value=\"".$row->_id."\"> ";
    echo "Comment: <input type=\"text\" id=\"text\" name=\"text\"> <input type=\"submit\" value=\"comment\">";
    echo "</form>";
}

?>

<html>
<head>
<link href="bootstrap-3.3-2.2-dist/css/bootstrap.min.css" rel="stylesheet">
</head>

<body>
<h2>Blog</h2>

<form action="post.php" method="post">
<p>Title: <input type="text" id="Title" name="Title"></p>
<p>Author: <input type="text" id="Author" name="Author"></p>
<p>Category: <input type="text" id="Category" name="Category"></p>
<p>Content: <textarea type="textarea" id="Content" name="Content"></textarea></p>
<p><input type="submit" value="submit"></p>
</form>
</body>
</html>