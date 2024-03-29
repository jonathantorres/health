<?php

use App\User;
use Illuminate\Database\Seeder;

class UsersTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        User::create([
            'name' => 'Jonathan',
            'last_name' => 'Torres',
            'email' => 'jonathantorres41@gmail.com',
            'password' => bcrypt('test'),
        ]);
    }
}
