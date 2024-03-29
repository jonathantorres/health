<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class AddUserIdToBloodPressuresTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('blood_pressures', function (Blueprint $table) {
            $table->unsignedInteger('user_id')->after('id')->default(1)->index();

            $table->foreign('user_id')->references('id')->on('users')->onDelete('cascade');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('blood_pressures', function (Blueprint $table) {
            $table->dropForeign('blood_pressures_user_id_foreign');
            $table->dropColumn('user_id');
        });
    }
}
