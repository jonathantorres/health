<?php

namespace App\Http\Controllers;

use App\Weight;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class WeightController extends Controller
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
     * Add a weight entry.
     * Shows the view and also process the form.
     *
     * @return Illuminate\View\View
     */
    public function add(Request $request)
    {
        if ($request->isMethod('post')) {
            $this->validate($request, [
                'weight' => 'required|numeric',
                'entered-date' => 'required',
            ]);

            Weight::create([
                'user_id' => Auth::user()->id,
                'weight' => $request->input('weight'),
                'entered_date' => $request->input('entered-date'),
            ]);

            return redirect()->route('add-weight')->with('success', 'Weight entry added succesfully.');
        }

        return view('weight.add', $this->data);
    }

    /**
     * View to show all weight entries.
     *
     * @return Illuminate\View\View
     */
    public function all()
    {
        $title = 'Weight Entries';
        $weightEntries = Auth::user()->weights()
                        ->orderBy('entered_date', 'desc')
                        ->paginate(20);

        $this->data['weightEntries'] = $weightEntries;
        $this->data['title'] = $title;

        return view('weight.all', $this->data);
    }

    public function edit(Request $request, $id)
    {
        $entry = Weight::find($id);

        // user doesn't own that weight entry
        if (empty($entry) || $entry->user_id !== Auth::user()->id) {
            return redirect()->route('index')
                             ->with('error', 'Error! You do not have access to that weight entry.');
        }

        if ($request->isMethod('post')) {
            $this->validate($request, [
                'weight' => 'required|numeric',
                'entered-date' => 'required',
            ]);

            $entry->weight = $request->input('weight');
            $entry->entered_date = $request->input('entered-date');
            $entry->save();

            $request->session()->flash('success', 'Weight entry updated succesfully.');
        }

        $this->data['entry'] = $entry;

        return view('weight.edit', $this->data);
    }
}
