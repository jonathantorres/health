@extends('layouts.app')

@section('content')
@include('partials.nav')

<div class="container">
    <div class="row">
        <div class="col-md-12">
            <div class="panel panel-default">
                <div class="panel-heading">Blood Pressure Readings</div>
                <div class="panel-body">
                    <div class="table-responsive">
                        <table class="table table-striped table-hover table-condensed">
                            <thead>
                                <tr>
                                    <th class="text-center">SYS</th>
                                    <th class="text-center">DIA</th>
                                    <th class="text-center">Pulse</th>
                                    <th class="text-center">Date</th>
                                    <th class="text-center">Severity</th>
                                    <th class="text-center">View Details</th>
                                </tr>
                            </thead>
                            <tbody>
                                @foreach($readings as $reading)
                                    <tr>
                                        <td class="text-center">{{ $reading->systolic }}</td>
                                        <td class="text-center">{{ $reading->diastolic }}</td>
                                        <td class="text-center">{{ $reading->pulse }}</td>
                                        <td class="text-center">{{ date('M, j Y', strtotime($reading->reading_date)) }}</td>
                                        <td class="text-center">Normal</td>
                                        <td class="text-center">
                                            <a class="" href="#">
                                                <span class="glyphicon glyphicon-search" aria-hidden="true"></span>
                                            </a>
                                        </td>
                                    </tr>
                                @endforeach
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
