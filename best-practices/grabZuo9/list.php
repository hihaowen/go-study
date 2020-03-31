<?php
/*****************************************************************
 * Copyright (C) 2020 Ltwen.com. All Rights Reserved.
 * Licensed http://www.apache.org/licenses/LICENSE-2.0
 * 文件名称：list.php
 * 创 建 者：hwz <haowen.hi@gmail.com>
 * 创建日期：2020年03月31日 00:30:59
 ****************************************************************/

$p = 0;
while($p++ < 5267) 
{
    // echo "curr page: ", $p, PHP_EOL;
    printListByPage($p); 
}

function printListByPage(int $page) {
        exec(sprintf("curl --retry 5 'https://www.zuo9.live/api/app/video/prefers?uuid=9dfee0e61d7911ea814256000274a59c&page=%d' -H 'authority: www.zuo9.live' -H 'accept: application/json, text/plain, */*' -H 'sec-fetch-dest: empty' -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36' -H 'token: d7caa50471b511eabff85a0002a03c65' -H 'sec-fetch-site: same-origin' -H 'sec-fetch-mode: cors' -H 'referer: https://www.zuo9.live/video/index?uuid=9dfee0e61d7911ea814256000274a59c&stack-key=94eced6f' -H 'accept-language: zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,ja;q=0.6'", $page), $output);

        $list = json_decode($output[0], true);
        $data = $list['result']['data'] ?? [];
        if (empty($data)) {
            return;
        }

        foreach($data as $perData) {
                echo $page, ",", $perData['uuid'], ",", $perData['title'], ",", $perData['thumbnail'], PHP_EOL;
        }
}

