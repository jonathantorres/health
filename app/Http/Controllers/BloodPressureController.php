<?php

namespace App\Http\Controllers;

use App\BloodPressure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

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
    public function add(Request $request)
    {
        if ($request->isMethod('post')) {
            $this->validate($request, [
                'sys' => 'required|numeric',
                'dia' => 'required|numeric',
                'pulse' => 'required|numeric',
                'reading-date' => 'required',
            ]);

            BloodPressure::create([
                'user_id' => Auth::user()->id,
                'systolic' => $request->input('sys'),
                'diastolic' => $request->input('dia'),
                'pulse' => $request->input('pulse'),
                'reading_date' => $request->input('reading-date'),
            ]);

            return redirect()->route('add-blood-pressure')->with('success', 'Blood Pressure reading added succesfully.');
        }

        return view('blood_pressure.add');
    }

    /**
     * View to show all blood pressure readings.
     *
     * @return Illuminate\View\View
     */
    public function all()
    {
        $title = 'Blood Pressure Readings';
        $readings = Auth::user()->bloodPressures()
                        ->orderBy('reading_date', 'desc')
                        ->paginate(20);

        return view('blood_pressure.all', [
            'readings' => $readings,
            'title' => $title,
        ]);
    }

    /**
     * Show details of a blood pressure reading.
     *
     * @param  int $id
     *
     * @return Illuminate\View\View
     */
    public function details($id)
    {
        $reading = BloodPressure::find($id);

        return view('blood_pressure.details', [
            'reading' => $reading,
        ]);
    }
}
