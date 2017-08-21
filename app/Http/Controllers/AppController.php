<?php

namespace App\Http\Controllers;

use App\BloodPressure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class AppController extends Controller
{
    /**
     * Create a new controller instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();

        $this->middleware('auth');
    }

    /**
     * Show the main dashboard view.
     *
     * @return Illuminate\View\View
     */
    public function index()
    {
        $title = 'Latest Blood Pressure Readings';
        $bloodPressureReadings = Auth::user()->bloodPressures()
                                     ->orderBy('reading_date', 'desc')
                                     ->limit(10)->get();

        $this->data['readings'] = $bloodPressureReadings;
        $this->data['title'] = $title;

        return view('index', $this->data);
    }
}
