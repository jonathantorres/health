@extends('layouts.app')

@section('content')
@include('partials.nav')

<div class="container">
    <div class="row">
        <div class="col-md-10 col-md-offset-1">
            @include('partials.flash_messages')

            <div class="panel panel-default">
                <div class="panel-heading">Add New Weight Entry</div>
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
                    <form class="form-horizontal" method="post" action="{{ route('add-weight') }}">
                        {{ csrf_field() }}
                        <div class="form-group">
                            <label for="weight" class="col-sm-2 control-label">Weight</label>
                            <div class="col-sm-7">
                                <input type="number" min="0" step="0.1" class="form-control" id="weight" name="weight" placeholder="Weight" required>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="entered-date" class="col-sm-2 control-label">Entry Date</label>
                            <div class="col-sm-7">
                                <input type="date" class="form-control" id="entered-date" name="entered-date" placeholder="Entry Date" required>
                            </div>
                        </div>
                        <div class="form-group">
                            <div class="col-sm-offset-2 col-sm-7">
                                <button type="submit" class="btn btn-primary">Add Weight Entry</button>
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
