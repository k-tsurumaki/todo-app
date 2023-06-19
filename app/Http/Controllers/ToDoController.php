<?php

namespace App\Http\Controllers;

use App\Http\Requests\ToDo\StoreRequest;
use App\Http\Requests\ToDo\UpdateRequest;
use App\Models\ToDo;
use App\Models\ToDoDetail;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

# コントローラーに書くのはモデルのどの機能を呼び出すのか（1、2行程度）
# それ以上の複雑な処理はモデルに記述する

class ToDoController extends Controller 
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // ToDoを取得する
        $toDos = ToDo::with('toDoDetails')->get(); # toDoDetailsも一緒に持ってくる

        // 取得したToDoを返却する
        return $toDos;
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(StoreRequest $request) # バリデーションのためにRequestをStoreRequestに変更（ここにバリデーション設定を書いてもいいが、コードが煩雑になる）
    {
        // 新規のToDoモデルを作成する
        $toDo = new ToDo(); # ToDoインスタンスを生成（Modelにあるのはモデルの設計図でこの行でそれに実体を与えている）

        // タイトルをToDoモデルに設定する
        $toDo->title = $request->get('title');

        // 空のToDoDetailを作成する
        $toDoDetail = new ToDoDetail();
        $toDoDetail->name = null;
        $toDoDetail->completed_flag = false;

        // DBにデータを登録する
        DB::transaction(function () use ($toDo, $toDoDetail){
            $toDo->save();
            $toDo->toDoDetails()->save($toDoDetail);
        });
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateRequest $request, $id)
    {
        // idに紐付くToDoモデルを取得する
        $toDo = ToDo::find($id);
        
        // タイトルをToDoモデルに設定する
        $toDo->title = $request->get('title');
        
        // ToDoデータベースを更新する
        $toDo -> save();
    }
    
    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        // idに紐付くToDoモデルを取得する
        $toDo = ToDo::find($id);
        
        // ToDoデータベースから対象のレコードを削除する
        $toDo -> delete();
    }
}
