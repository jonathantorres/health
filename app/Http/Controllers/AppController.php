<?php

namespace App\Http\Controllers;

use App\BloodPressure;
use Illuminate\Http\Request;

class AppController extends Controller
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
     * Show the main dashboard view.
     *
     * @return Illuminate\View\View
     */
    public function index()
    {
        $bloodPressureReadings = BloodPressure::orderBy('reading_date', 'desc')->limit(10)->get();

        return view('index', [
            'readings' => $bloodPressureReadings,
        ]);
    }
}
