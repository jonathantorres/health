@extends('layouts.app')

@section('content')
@include('partials.nav')

<div class="container">
    <div class="row">
        <div class="col-md-12">
            @include('partials.flash_messages')
        </div>
    </div>

    <div class="row">
        <div class="col-md-12">
            @include('partials.blood_pressure.readings', ['title' => $pressureTitle])
        </div>
    </div>
    <div class="row">
        <div class="col-md-12">
            @include('partials.weight.entries', ['title' => $weightTitle])
        </div>
    </div>
    <div class="row hidden">
        <div class="col-md-12">
            <div class="panel panel-default">
                <div class="panel-heading">Consumed Food</div>
                <div class="panel-body">
                    <div class="alert alert-warning">
                        There are no consumed food entries.
                    </div>
                </div>
            </div>
        </div>
    </div>
    @include('partials.footer')
</div>
@endsection
