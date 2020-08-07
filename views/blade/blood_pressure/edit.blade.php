@extends('layouts.app')

@section('content')
@include('partials.nav')

<div class="container">
    <div class="row">
        <div class="col-md-10 col-md-offset-1">
            @include('partials.flash_messages')

            <div class="panel panel-default">
                <div class="panel-heading">Update Blood Pressure Reading</div>
                <div class="panel-body">
                    @if (count($errors) > 0)
                        <div class="alert alert-danger">
                            <ul>
                                @foreach ($errors->all() as $error)
                                    <li>{{ $error }}</li>
                                @endforeach
                            </ul>
                        </div>
                    @endif
                    <form class="form-horizontal" method="post" action="{{ route('edit-blood-pressure', ['id' => $reading->id]) }}">
                        {{ csrf_field() }}
                        <div class="form-group">
                            <label for="sys" class="col-sm-2 control-label">SYS</label>
                            <div class="col-sm-7">
                                <input type="number" class="form-control" id="sys" name="sys" placeholder="SYS" value="{{ $reading->systolic }}" required>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="dia" class="col-sm-2 control-label">DIA</label>
                            <div class="col-sm-7">
                                <input type="number" class="form-control" id="dia" name="dia" placeholder="DIA" value="{{ $reading->diastolic }}" required>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="pulse" class="col-sm-2 control-label">Pulse</label>
                            <div class="col-sm-7">
                                <input type="number" class="form-control" id="pulse" name="pulse" placeholder="Pulse" value="{{ $reading->pulse }}" required>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="reading-date" class="col-sm-2 control-label">Reading Date</label>
                            <div class="col-sm-7">
                                <input type="date" class="form-control" id="reading-date" name="reading-date" placeholder="Reading Date" value="{{ $reading->reading_date->toDateString() }}" required>
                            </div>
                        </div>
                        <div class="form-group">
                            <div class="col-sm-offset-2 col-sm-7">
                                <button type="submit" class="btn btn-primary">Update Reading</button>
                                <a href="{{ route('index') }}" class="btn btn-link">Cancel</a>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
    @include('partials.footer')
</div>
@endsection
