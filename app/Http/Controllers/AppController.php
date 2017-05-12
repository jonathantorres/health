<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class AppController extends Controller
{
    /**
     * Show the main index view.
     *
     * @return Illuminate\View\View
     */
    public function index()
    {
        return view('index');
    }
}
