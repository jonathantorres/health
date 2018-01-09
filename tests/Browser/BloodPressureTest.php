<?php

namespace Tests\Browser;

use App\BloodPressure;
use App\User;
use Illuminate\Foundation\Testing\DatabaseMigrations;
use Laravel\Dusk\Browser;
use Tests\DuskTestCase;

class BloodPressureTest extends DuskTestCase
{
    use DatabaseMigrations;

    public function setUp()
    {
        parent::setUp();
    }

     /** @test */
    public function user_should_see_all_blood_pressure_readings()
    {
        $user = factory(User::class)->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Blood Pressure Readings')
                    ->clickLink('See all readings')
                    ->assertPathIs('/blood-pressure/all')
                    ->assertSee('Blood Pressure Readings')
                    ->logout();
        });
    }

    /** @test */
    public function user_can_see_details_of_a_reading_from_main_dashboard()
    {
        $user = factory(User::class)->create();
        $readings = factory(BloodPressure::class, 5)->create([
            'user_id' => $user->id,
        ]);
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Blood Pressure Readings')
                    ->click('.glyphicon.glyphicon-search')
                    ->assertSee('Reading Details')
                    ->logout();
        });
    }

    /** @test */
    public function user_can_see_details_of_a_reading_from_all_readings_view()
    {
        $user = factory(User::class)->create();
        $readings = factory(BloodPressure::class, 5)->create([
            'user_id' => $user->id,
        ]);
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Blood Pressure Readings')
                    ->clickLink('See all readings')
                    ->click('.glyphicon.glyphicon-search')
                    ->assertSee('Reading Details')
                    ->logout();
        });
    }

    /** @test */
    public function user_can_enter_a_new_blood_pressure_reading()
    {
        $user = factory(User::class)->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Blood Pressure Readings')
                    ->clickLink('Enter new reading')
                    ->assertSee('Add New Blood Pressure Reading')
                    ->type('sys', 120)
                    ->type('dia', 80)
                    ->type('pulse', 75)
                    // ->type('reading-date', '09/12/2017')
                    ->script([
                        "document.querySelector('#reading-date').value = '2017-09-12'",
                    ]);

            $browser->press('Add Reading')
                    ->assertPathIs('/blood-pressure/add')
                    ->assertSee('Blood Pressure reading added succesfully.')
                    ->logout();
        });
    }

    /** @test */
    public function user_can_edit_an_existing_blood_pressure_reading()
    {
        $user = factory(User::class)->create();
        $readings = factory(BloodPressure::class, 5)->create([
            'user_id' => $user->id,
        ]);
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Blood Pressure Readings')
                    ->clickLink('See all readings')
                    ->click('.glyphicon.glyphicon-edit')
                    ->assertSee('Update Blood Pressure Reading')
                    ->type('sys', 120)
                    ->type('dia', 80)
                    ->type('pulse', 75)
                    ->script([
                        "document.querySelector('#reading-date').value = '2017-09-12'",
                    ]);

            $browser->press('Update Reading')
                    ->assertPathIs('/blood-pressure/edit/1')
                    ->assertSee('Blood Pressure reading updated succesfully.')
                    ->logout();
        });
    }

    /** @test */
    public function user_can_delete_an_existing_blood_pressure_reading()
    {
        $user = factory(User::class)->create();
        $readings = factory(BloodPressure::class, 5)->create([
            'user_id' => $user->id,
        ]);
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Blood Pressure Readings')
                    ->clickLink('See all readings')
                    ->click('.glyphicon.glyphicon-remove')
                    ->waitForText('Are you sure you wish to delete this reading?')
                    ->press('Yes')
                    ->assertPathIs('/')
                    ->assertSee('Blood Pressure reading deleted succesfully.')
                    ->logout();
        });
    }
}
