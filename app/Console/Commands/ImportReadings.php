<?php

namespace App\Console\Commands;

use Carbon\Carbon;
use Illuminate\Console\Command;
use Illuminate\Support\Facades\DB;

class ImportReadings extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'readings:import';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Import initial blood pressure readings.';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $this->info('Importing predefined readings...');

        DB::table('blood_pressures')->insert($this->readingsToImport());

        $this->info('Done.');
    }

    /**
     * Define and return the readings to import.
     *
     * @todo  Add the actual readings.
     *
     * @return array
     */
    private function readingsToImport()
    {
        return [
            // [
            //     'systolic' => 123,
            //     'diastolic' => 81,
            //     'pulse' => 75,
            //     'reading_date' => Carbon::createFromDate(2016, 6, 12),
            //     'created_at' => Carbon::now(),
            //     'updated_at' => Carbon::now(),
            // ],
        ];
    }
}
