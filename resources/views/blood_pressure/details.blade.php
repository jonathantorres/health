@extends('layouts.app')

@section('content')
@include('partials.nav')

<div class="container">
    <div class="row">
        <div class="col-md-10 col-md-offset-1">
            <h2>Reading Details</h2>
            <hr>
            <p>Systolic: <strong>{{ $reading->systolic }}</strong></p>
            <p>Diastolic: <strong>{{ $reading->diastolic }}</strong></p>
            <p>Pulse: <strong>{{ $reading->pulse }}</strong></p>
            <p>Date: <strong>{{ date('M, j Y', strtotime($reading->reading_date)) }}</strong></p>
            <p class="text-{{ $reading->severity()['class'] }}">
                Severity: <strong>{{ $reading->severity()['text'] }}</strong>
            </p>
            <hr>
            <a href="{{ url()->previous() }}" class="btn btn-default">Go Back</a>
        </div>
    </div>
    @include('partials.footer')
</div>
@endsection
