<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class ToDo extends Model # モデルはToDosの一行と紐づいている
{
    use HasFactory;

    public function toDoDetails()
    {
        return $this->hasMany(ToDoDetail::class); # ToDoはToDoDetailをたくさん持っているというリレーションを意味している（このToDoDetailはモデル名を指している）
    }

    public function delete()
    {
        // 関連するToDoDetailsのレコードを削除する
        $this->toDoDetails()->delete();

        // ToDoのレコードを削除する
        return parent::delete();
    }
}
