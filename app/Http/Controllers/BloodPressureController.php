<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class BloodPressureController extends Controller
{
    /**
     * Create a new controller instance.
     *
     * @return void
     */
    public function __construct()
    {
        $this->middleware('auth');
    }

    /**
     * Add a new blood pressure reading.
     * Shows the view and also process the form.
     *
     * @return Illuminate\View\View
     */
    public function add()
    {
        return view('blood_pressure.add');
    }
}
