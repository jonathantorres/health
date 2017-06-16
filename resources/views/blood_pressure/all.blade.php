@extends('layouts.app')

@section('content')
@include('partials.nav')

<div class="container">
    <div class="row">
        <div class="col-md-12">
            @include('partials.blood_pressure.readings')
        </div>
    </div>
    <div class="row">
        <div class="col-md-12 text-center">
            {{ $readings->links() }}
        </div>
    </div>
</div>
@endsection
