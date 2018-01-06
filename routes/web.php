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
Route::any('blood-pressure/edit/{id}', 'BloodPressureController@edit')->name('edit-blood-pressure');
Route::get('blood-pressure/delete/{id}', 'BloodPressureController@delete')->name('delete-blood-pressure');
Route::any('weight/add', 'WeightController@add')->name('add-weight');
Route::get('weight/all', 'WeightController@all')->name('all-weights');
Route::any('weight/edit/{id}', 'WeightController@edit')->name('edit-weight');
Route::get('weight/delete/{id}', 'WeightController@delete')->name('delete-weight');

Auth::routes();
