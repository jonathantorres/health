<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', 'AppController@index')->name('index');
Route::any('blood-pressure/add', 'BloodPressureController@add')->name('add-blood-pressure');
Route::get('blood-pressure/all', 'BloodPressureController@all')->name('all-blood-pressure');
Route::get('blood-pressure/details/{id}', 'BloodPressureController@details')->name('blood-pressure-details');

Auth::routes();
