@extends('layouts.app')

@section('content')
@include('partials.nav')

<div class="container">
    <div class="row">
        <div class="col-md-12">
            @include('partials.blood_pressure.readings', [ 'title' => $title])
        </div>
    </div>
    <div class="row">
        <div class="col-md-12">
            <div class="panel panel-default">
                <div class="panel-heading">Weight</div>
                <div class="panel-body">
                    <div class="alert alert-warning">
                        There are no weight entries.
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="row">
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
</div>
@endsection
