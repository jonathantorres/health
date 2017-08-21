<?php

namespace App\Http\Controllers;

use App\Version;
use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Routing\Controller as BaseController;

class Controller extends BaseController
{
    use AuthorizesRequests, DispatchesJobs, ValidatesRequests;

    /**
     * The data that is going to be passed to the view.
     *
     * @var array
     */
    protected $data = [];

    public function __construct()
    {
        $this->data['version'] = Version::VERSION;
    }
}
