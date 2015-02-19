#GoTableThis
This Go package will allow you to format and print tables with data.  The package is flexible as far as the ways the data can be arranged and the respective output formats.

## Overview
In order to print structured information from Golang there are many options to pursue.  One option may to leverage a structured medium like XML, JSON, or YAML. The benefit of these methods would that they are readily able to be manipulated programmatically and able to print any length of data predictably.

Otherwise you may want a legacy look and feel, and in this case it may be ideal to format the output in a "pretty table" or "pretty list" with columns and rows and proper divisors.  

The actual table formatting aspect is performed from the following  [TableWriter](github.com/olekukonko/tablewriter) package.

## Tables
There are three direct methods that receive from a ```Table``` struct.  These methods format the data slightly differently.

Included with the package is the ```gotablethis_test.go``` file which can be ran directly with a ```go test``` to view the output methods.

Notice how we are sending a reflect Value since this allows us to receive any struct.

### PrintColumn
This is the simplest method method where it expects only to see a column name and an array of something that can be converted to strings.

    testing := []string{"test1"}

    table := Table{
      Header:  []string{"testheader1"},
      RowData: reflect.ValueOf(&testing).Elem(),
    }

    table.PrintColumn()

This outputs the following table.

    +-------------+
    | TESTHEADER1 |
    +-------------+
    | test1       |
    +-------------+

### PrintKeyValueTable
This method expects that we are sending a single struct with parameters that can be converted to strings.  The output is a single table with two columns of key and value.  

    testing := teststruct{
      Test1: "first",
      Test2: 1,
      Test3: true,
    }

    table := Table{
      RowData: reflect.ValueOf(&testing).Elem(),
    }

    table.PrintKeyValueTable()

The output from this method looks like the following.

    +-------+-------+
    |  KEY  | VALUE |
    +-------+-------+
    | Test1 | first |
    | Test2 | 1     |
    | Test3 | true  |
    +-------+-------+

### PrintTable
In the case that you would like to send a struct with an array of internal structs and choose specific headers you would use ```PrintTable```.

Notice how there is a ```Header``` parameter being specified.  These names can be arbitrary and map to the ```Columns``` data based on ordered position.  The ```Columns``` array of strings needs to match in a case-sensitive way the parameters being received from the structs.

    testing := teststructs{
      teststruct{
        Test1: "first",
        Test2: 1,
        Test3: true,
      },
      teststruct{
        Test1: "second",
        Test2: 2,
        Test3: false,
      },
    }

    table := Table{
      Header:  []string{"test1", "test2", "test3"},
      Columns: []string{"Test1", "Test2", "Test3"},
      RowData: reflect.ValueOf(&testing).Elem(),
    }

    table.PrintTable()





Licensing
---------
Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

Support
-------
Please file bugs and issues at the Github issues page. For more general discussions you can contact the EMC Code team at <a href="https://groups.google.com/forum/#!forum/emccode-users">Google Groups</a> or tagged with **EMC** on <a href="https://stackoverflow.com">Stackoverflow.com</a>. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
