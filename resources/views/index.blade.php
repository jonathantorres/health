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
        <div class="col-md-12">
            <div class="panel panel-default">
                <div class="panel-heading">Weight</div>
                <div class="panel-body">
                    <div class="table-responsive">
                        <table class="table table-striped table-hover table-condensed">
                            <thead>
                                <tr>
                                    <th class="text-center">One</th>
                                    <th class="text-center">Two</th>
                                    <th class="text-center">Three</th>
                                </tr>
                            </thead>
                            <tbody>
                                @for($i = 0; $i < 10; $i++)
                                    <tr>
                                        <td class="text-center">Weight 1</td>
                                        <td class="text-center">Weight 2</td>
                                        <td class="text-center">Weight 3</td>
                                    </tr>
                                @endfor
                            </tbody>
                        </table>
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
                    <div class="table-responsive">
                        <table class="table table-striped table-hover table-condensed">
                            <thead>
                                <tr>
                                    <th class="text-center">One</th>
                                    <th class="text-center">Two</th>
                                    <th class="text-center">Three</th>
                                </tr>
                            </thead>
                            <tbody>
                                @for($i = 0; $i < 10; $i++)
                                    <tr>
                                        <td class="text-center">Food 1</td>
                                        <td class="text-center">Food 2</td>
                                        <td class="text-center">Food 3</td>
                                    </tr>
                                @endfor
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
