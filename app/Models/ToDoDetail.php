<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class ToDoDetail extends Model # テーブルにアクセスするためのクラスをモデルという
{
    use HasFactory;

    public function todo()
    {
        return $this->belongsTo(ToDo::class); # ToDoDetailは一つのToDoに属しているというリレーションを意味している
    }

    public function getCompletedFlagAttribute($value)
    {
        return $value == 1;
    }
}
