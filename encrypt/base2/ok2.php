<?php

     function encrypt3DES($input, $key)
    {
        $size = mcrypt_get_block_size("tripledes", "ecb");
        $input = pkcs5Pad($input, $size);
        $key = str_pad($key, 24, '0');
        $td = mcrypt_module_open("tripledes", '', "ecb", '');
        $iv = @mcrypt_create_iv(mcrypt_enc_get_iv_size($td), 2);
        @mcrypt_generic_init($td, $key, $iv);
        $data = mcrypt_generic($td, $input);
        mcrypt_generic_deinit($td);
        mcrypt_module_close($td);
        $data = base64_encode($data);
        return $data;
    }
       function pkcs5Pad($text, $blocksize)
    {
        $pad = $blocksize - (strlen($text) % $blocksize);
        return $text . str_repeat(chr($pad), $pad);
    }

     function pkcs5Unpad($text)
    {
        $pad = ord($text{strlen($text) - 1});
        if ($pad > strlen($text)) {
            return false;
        }
        if (strspn($text, chr($pad), strlen($text) - $pad) != $pad) {
            return false;
        }
        return substr($text, 0, -1 * $pad);
    }

     $name = encrypt3DES('路雪峰', 'a677c601941a4cfe9f1ab63432259f5f');
     var_dump( $name);