<?php

use App\Http\Controllers\AppController;
use App\Http\Controllers\BloodPressureController;
use App\Http\Controllers\WeightController;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Route;

Route::get('/', [AppController::class, 'index'])->name('index');

Route::match(['get', 'post'], 'blood-pressure/add', [BloodPressureController::class, 'add'])->name('add-blood-pressure');
Route::get('blood-pressure/all', [BloodPressureController::class, 'all'])->name('all-blood-pressure');
Route::get('blood-pressure/details/{id}', [BloodPressureController::class, 'details'])->name('blood-pressure-details');
Route::match(['get', 'post'], 'blood-pressure/edit/{id}', [BloodPressureController::class, 'edit'])->name('edit-blood-pressure');
Route::get('blood-pressure/delete/{id}', [BloodPressureController::class, 'delete'])->name('delete-blood-pressure');

Route::match(['get', 'post'], 'weight/add', [WeightController::class, 'add'])->name('add-weight');
Route::get('weight/all', [WeightController::class, 'all'])->name('all-weights');
Route::match(['get', 'post'], 'weight/edit/{id}', [WeightController::class, 'edit'])->name('edit-weight');
Route::get('weight/delete/{id}', [WeightController::class, 'delete'])->name('delete-weight');

Auth::routes();

